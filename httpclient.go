package moysklad

import (
	"encoding/json"
	"github.com/tirprox/moysklad/codec"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Result struct {
	Index int
	Res   *http.Response
	Err   error
}

type Response struct {
	Index int
	Body  []byte
}

type GenericApiResponse struct {
	Meta codec.Meta `json:"meta"`
}

func (c moyskladClient) boundedParallelGet(urls []string, concurrencyLimit int) []Result {

	semaphoreChan := make(chan struct{}, concurrencyLimit)

	resultsChan := make(chan *Result)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for i, generatedUrl := range urls {

		go func(i int, url string) {

			semaphoreChan <- struct{}{}
			res, err := c.MakeRequest(url, "GET")
			result := &Result{i, res, err}
			resultsChan <- result
			<-semaphoreChan

		}(i, generatedUrl)
	}

	var results []Result

	for {
		result := <-resultsChan
		results = append(results, *result)

		if len(results) == len(urls) {
			break
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	return results
}

func (c moyskladClient) GetAll(url string) []interface{} {

	res, err := c.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	firstResponse := Response{}

	firstResponse.Index = 0

	firstResponse.Body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	responses := []Response{}
	responses = append(responses, firstResponse)

	meta := DecodeMeta(firstResponse.Body)
	urls := makeUrlList(url, meta.Size)

	if len(urls) > 0 {
		results := c.boundedParallelGet(urls, 5)
		for _, result := range results {
			index := result.Index + 1
			body, err := ioutil.ReadAll(result.Res.Body)
			if err != nil {
				log.Fatal(err)
			}

			response := Response{index, body}
			responses = append(responses, response)
			defer result.Res.Body.Close()
		}

	}

	var data []interface{}

	for _, response := range responses {
		var result map[string]interface{}
		json.Unmarshal(response.Body, &result)
		rows := result["rows"].([]interface{})
		for _, row := range rows {
			data = append(data, row)
		}

	}

	return data
}

func (c moyskladClient) Get(url string) (*http.Response, error) {
	res, err := c.MakeRequest(url, "GET")
	return res, err
}

var client = &http.Client{Timeout: 60 * time.Second}

func (c moyskladClient) MakeRequest(url string, method string) (res *http.Response, err error) {
	req, error1 := http.NewRequest(method, url, nil)
	if error1 != nil {
		log.Fatal(err)
	}

	c.setHeaders(req)

	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func (c moyskladClient) setHeaders(req *http.Request) {
	req.SetBasicAuth(c.login, c.password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

func DecodeMeta(responseBody []byte) *codec.Meta {
	response := GenericApiResponse{}
	json.Unmarshal(responseBody, &response)
	return &response.Meta
}

func makeUrlList(baseUrl string, size int) (urls []string) {
	offset := 100
	limit := 100

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	for offset < size {
		query := u.Query()
		query.Set("offset", strconv.Itoa(offset))
		u.RawQuery = query.Encode()
		urls = append(urls, u.String())
		offset += limit
	}
	return

}

func makeEntityUrl(entityName string) (url string) {
	url = codec.API_BASE + "entity/" + entityName + "/"
	return
}

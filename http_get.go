package moysklad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func (c moyskladClient) GetOne(entity string, id string, params QueryParams) map[string]interface{} {

	url := params.Url(makeEntityUrl(entity) + id)

	result := c.GetRequest(url)

	return result
}

func (c moyskladClient) GetMany(entity string, params QueryParams) []map[string]interface{} {

	url := params.Url(makeEntityUrl(entity))

	result := c.GetRequest(url)

	rows := extractRows(result)
	meta := extractMeta(result)

	size := int(meta["size"].(float64))

	if size > 100 {
		urls := makeUrlList(url, size)
		results := c.concurrentGet(urls, 5)

		for _, result := range results {
			extracted := extractRows(result.Response)
			rows = append(rows, extracted...)
		}
	}

	return rows
}

func (c moyskladClient) GetRequest(url string) (result map[string]interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	c.setHeaders(req)

	fmt.Println("Requesting: GET ", url)

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (c moyskladClient) concurrentGet(urls []string, concurrencyLimit int) []ResponseHolder {

	semaphoreChannel := make(chan struct{}, concurrencyLimit)
	responseChannel := make(chan *ResponseHolder)

	defer func() {
		close(semaphoreChannel)
		close(responseChannel)
	}()

	for i, genUrl := range urls {

		go func(i int, url string) {

			semaphoreChannel <- struct{}{}

			response := c.GetRequest(url)
			responseHolder := &ResponseHolder{i, response}
			responseChannel <- responseHolder

			<-semaphoreChannel

		}(i, genUrl)
	}

	var results []ResponseHolder

	for {
		result := <-responseChannel
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

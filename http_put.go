package moysklad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func (c moyskladClient) PutRequest(url string, data interface{}) (result map[string]interface{}) {

	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	//err = ioutil.WriteFile("request.json", requestBody, 0644)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	c.setHeaders(request)

	fmt.Println("Requesting: PUT ", url)

	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (c moyskladClient) Put(entity string, id string, data interface{}) map[string]interface{} {

	url := makeEntityUrl(entity) + id

	result := c.PutRequest(url, data)

	return result
}

func (c moyskladClient) concurrentPut(urls []string, concurrencyLimit int) []ResponseHolder {

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

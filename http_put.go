package moysklad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	fmt.Println("Requesting: POST ", url)

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

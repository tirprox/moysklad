package moysklad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//TODO add filter and parameters (forwarded to query constructor as a map of string -> string)

func (c moyskladClient) GetOne(entity string, id string, params QueryParams) map[string]interface{} {

	url := params.Url(makeEntityUrl(entity) + id)
	//url := makeEntityUrl(entity) + id
	//fmt.Println(url)

	res, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	meta := result["meta"].(map[string]interface{})
	fmt.Println(meta["type"].(string))

	return result
}

func (c moyskladClient) GetMany(entity string, params QueryParams) map[string]interface{} {

	url := params.Url(makeEntityUrl(entity))
	//url := makeEntityUrl(entity)

	res, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	//process first response

	rows := extractRows(result)
	meta := extractMeta(result)

	fmt.Println(meta["size"].(float64))

	for _, row := range rows {
		fmt.Println(row["name"].(string))
	}

	return result
}

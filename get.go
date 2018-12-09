package moysklad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (c moyskladClient) GetOne(entity string, id string) map[string]interface{} {
	url := makeEntityUrl(entity) + id
	fmt.Println(url)

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

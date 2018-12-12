package moysklad

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var httpClient = &http.Client{Timeout: 30 * time.Second}

func (c moyskladClient) setHeaders(req *http.Request) {
	req.SetBasicAuth(c.login, c.password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

func makeEntityUrl(entityName string) (url string) {
	url = API_BASE + "entity/" + entityName + "/"
	return
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

type ResponseHolder struct {
	Index    int
	Response map[string]interface{}
}

func extractRows(response map[string]interface{}) []map[string]interface{} {

	if response["rows"] != nil {
		rows := response["rows"].([]interface{})

		var rowMap []map[string]interface{}

		for _, row := range rows {
			rmap := row.(map[string]interface{})
			rowMap = append(rowMap, rmap)
		}

		return rowMap
	} else {
		return []map[string]interface{}{response}
	}

}

func extractMeta(response map[string]interface{}) map[string]interface{} {
	if response["meta"] != nil {
		meta := response["meta"].(map[string]interface{})
		return meta

	} else {
		return nil
	}

}

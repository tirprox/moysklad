package moysklad

import (
	"github.com/mitchellh/mapstructure"
	"github.com/tirprox/moysklad/codec"
)

type GenericResponse struct {
	Meta codec.Meta
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
		var rMap []map[string]interface{}
		rMap = append(rMap, response)
		return rMap

	}

}

func extractMeta(response map[string]interface{}) map[string]interface{} {
	meta := response["meta"].(map[string]interface{})
	return meta
}

func ToEntity(row map[string]interface{}) (item interface{}, entityType string) {

	meta := row["meta"].(map[string]interface{})
	entityType = meta["type"].(string)

	//fmt.Println(entityType)

	switch entityType {

	/*case "product":
		result := Product
		return
	case "variant":

	case "counterparty":

	case "customerorder":*/

	case "store":
		result := codec.Store{}
		mapstructure.Decode(row, &result)
		return result, entityType

	default:
		result := GenericResponse{}
		mapstructure.Decode(row, &result)
		return result, entityType
	}

}

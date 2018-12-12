package moysklad

import (
	"log"
	"net/url"
)

type QueryParams struct {
	Filters []Filter
	Query   map[string]string
}

type Filter struct {
	Key      string
	Value    string
	Operator string
}

type Query struct {
	Entity string
	Params QueryParams
}

func (params QueryParams) GetFilterValue() (filterStr string) {
	filterStr = ""
	for _, filter := range params.Filters {
		tmp := filter.Key + filter.Operator + filter.Value + ";"
		filterStr = filterStr + tmp
	}

	return
}

func (params *QueryParams) AddFilter(key string, value string, operator string) {
	params.Filters = append(params.Filters, Filter{key, value, operator})
}

func (params *QueryParams) AddQuery(key string, value string) {
	if params.Query == nil {
		params.Query = make(map[string]string)
	}

	params.Query[key] = value
}

func (params QueryParams) Url(baseUrl string) (encodedUrl string) {

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	query := u.Query()

	filterValue := params.GetFilterValue()
	if filterValue != "" {
		query.Set("filter", filterValue)
	}

	for key, value := range params.Query {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	encodedUrl = u.String()
	return
}

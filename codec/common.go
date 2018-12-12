package codec

type Meta struct {
	Href         string `json:"href,omitempty"`
	MetadataHref string `json:"metadataHref,omitempty"`
	Type         string `json:"type,omitempty"`
	MediaType    string `json:"mediaType,omitempty"`
	Size         int    `json:"size,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
	NextHref     string `json:"nextHref,omitempty"`
	UUIDHref     string `json:"uuidHref,omitempty"`
}

type Context struct {
	Employee struct {
		Meta *Meta `json:"meta,omitempty"`
	} `json:"employee,omitempty"`
}

type Group struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Owner struct {
	Meta *Meta `json:"meta,omitempty"`
}

type SalePrices []struct {
	Value     float64   `json:"value,omitempty"`
	Currency  *Currency `json:"currency,omitempty"`
	PriceType string    `json:"priceType,omitempty"`
}

type Currency struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Rate struct {
	Currency *Currency `json:"currency,omitempty"`
}

type BuyPrice struct {
	Value    float64   `json:"value,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
}

type Characteristics []struct {
	Meta  *Meta  `json:"meta,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Uom struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Stock struct {
	Stock map[string]string `json:"stock,omitempty"`
}

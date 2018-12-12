package codec

type Assortment struct {
	Meta               *Meta            `json:"meta,omitempty"`
	ID                 string           `json:"id,omitempty"`
	AccountID          string           `json:"accountId,omitempty"`
	Owner              *Owner           `json:"owner,omitempty"`
	Shared             bool             `json:"shared,omitempty"`
	Group              *Group           `json:"group,omitempty"`
	Version            int              `json:"version,omitempty"`
	Updated            string           `json:"updated,omitempty"`
	Name               string           `json:"name,omitempty"`
	Code               string           `json:"code,omitempty"`
	ExternalCode       string           `json:"externalCode,omitempty"`
	Archived           bool             `json:"archived,omitempty"`
	PathName           string           `json:"pathName,omitempty"`
	Uom                Uom              `json:"uom,omitempty"`
	MinPrice           float64          `json:"minPrice,omitempty"`
	SalePrices         *SalePrices      `json:"salePrices,omitempty"`
	BuyPrice           *BuyPrice        `json:"buyPrice,omitempty"`
	Article            string           `json:"article,omitempty"`
	Weight             float64          `json:"weight,omitempty"`
	Volume             float64          `json:"volume,omitempty"`
	Barcodes           []string         `json:"barcodes,omitempty"`
	ModificationsCount int              `json:"modificationsCount,omitempty"`
	IsSerialTrackable  bool             `json:"isSerialTrackable,omitempty"`
	Stock              float64          `json:"stock,omitempty"`
	Reserve            float64          `json:"reserve,omitempty"`
	InTransit          float64          `json:"inTransit,omitempty"`
	Quantity           float64          `json:"quantity,omitempty"`
	Characteristics    *Characteristics `json:"characteristics,omitempty"`
	Product            *Product         `json:"product,omitempty"`
}

package codec

type Customerorder struct {
	Meta                *Meta                `json:"meta,omitempty"`
	ID                  string               `json:"id,omitempty"`
	AccountID           string               `json:"accountId,omitempty"`
	Owner               *Owner               `json:"owner,omitempty"`
	Shared              bool                 `json:"shared,omitempty"`
	Group               *Group               `json:"group,omitempty"`
	Version             int                  `json:"version,omitempty"`
	Updated             string               `json:"updated,omitempty"`
	Name                string               `json:"name,omitempty"`
	ExternalCode        string               `json:"externalCode,omitempty"`
	Moment              string               `json:"moment,omitempty"`
	Applicable          bool                 `json:"applicable,omitempty"`
	Rate                *Rate                `json:"rate,omitempty"`
	Sum                 int                  `json:"sum,omitempty"`
	Store               *Store               `json:"store,omitempty"`
	Agent               *Agent               `json:"agent,omitempty"`
	Organization        *Organization        `json:"organization,omitempty"`
	OrganizationAccount *OrganizationAccount `json:"organizationAccount,omitempty"`
	State               *State               `json:"state,omitempty"`
	Documents           *Documents           `json:"documents,omitempty"`
	Created             string               `json:"created,omitempty"`
	Positions           *Positions           `json:"positions,omitempty"`
	VatEnabled          bool                 `json:"vatEnabled,omitempty"`
	VatIncluded         bool                 `json:"vatIncluded,omitempty"`
	VatSum              int                  `json:"vatSum,omitempty"`
	PayedSum            int                  `json:"payedSum,omitempty"`
	ShippedSum          int                  `json:"shippedSum,omitempty"`
	InvoicedSum         int                  `json:"invoicedSum,omitempty"`
	ReservedSum         int                  `json:"reservedSum,omitempty"`
	Demands             []struct {
		Meta *Meta `json:"meta,omitempty"`
	} `json:"demands,omitempty"`
}

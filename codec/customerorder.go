package codec

type Customerorder struct {
	Meta Meta `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     Owner `json:"owner"`
	Shared bool `json:"shared"`
	Group  Group `json:"group"`
	Version      int    `json:"version"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Moment       string `json:"moment"`
	Applicable   bool   `json:"applicable"`
	Rate         struct {
		Currency struct {
			Meta `json:"meta"`
		} `json:"currency"`
	} `json:"rate"`
	Sum   int `json:"sum"`
	Store Store `json:"store"`
	Agent struct {
		Meta Meta `json:"meta"`
		ID        string `json:"id"`
		AccountID string `json:"accountId"`
		Owner     struct {
			Meta struct {
				Href         string `json:"href"`
				MetadataHref string `json:"metadataHref"`
				Type         string `json:"type"`
				MediaType    string `json:"mediaType"`
				UUIDHref     string `json:"uuidHref"`
			} `json:"meta"`
		} `json:"owner"`
		Shared bool `json:"shared"`
		Group  struct {
			Meta struct {
				Href         string `json:"href"`
				MetadataHref string `json:"metadataHref"`
				Type         string `json:"type"`
				MediaType    string `json:"mediaType"`
			} `json:"meta"`
		} `json:"group"`
		Version      int    `json:"version"`
		Updated      string `json:"updated"`
		Name         string `json:"name"`
		ExternalCode string `json:"externalCode"`
		Archived     bool   `json:"archived"`
		Created      string `json:"created"`
		CompanyType  string `json:"companyType"`
		Phone        string `json:"phone"`
		Accounts     struct {
			Meta Meta `json:"meta"`
		} `json:"accounts"`
		Tags           []string `json:"tags"`
		Contactpersons struct {
			Meta Meta `json:"meta"`
		} `json:"contactpersons"`
		Notes struct {
			Meta Meta `json:"meta"`
		} `json:"notes"`
		State struct {
			Meta Meta `json:"meta"`
		} `json:"state"`
		SalesAmount int `json:"salesAmount"`
	} `json:"agent"`
	Organization struct {
		Meta Meta `json:"meta"`
	} `json:"organization"`
	OrganizationAccount struct {
		Meta Meta `json:"meta"`
	} `json:"organizationAccount"`
	State struct {
		Meta Meta `json:"meta"`
	} `json:"state"`
	Documents struct {
		Meta Meta `json:"meta"`
	} `json:"documents"`
	Created   string `json:"created"`
	Positions struct {
		Meta Meta `json:"meta"`
	} `json:"positions"`
	VatEnabled  bool `json:"vatEnabled"`
	VatIncluded bool `json:"vatIncluded"`
	VatSum      int  `json:"vatSum"`
	PayedSum    int  `json:"payedSum"`
	ShippedSum  int  `json:"shippedSum"`
	InvoicedSum int  `json:"invoicedSum"`
	ReservedSum int  `json:"reservedSum"`
	Demands     []struct {
		Meta Meta `json:"meta"`
	} `json:"demands"`
}

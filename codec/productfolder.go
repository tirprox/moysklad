package codec

type ProductFolder struct {
	Meta          *Meta  `json:"meta,omitempty"`
	ID            string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
	Owner         *Owner `json:"owner,omitempty"`
	Shared        bool   `json:"shared,omitempty"`
	Group         *Group `json:"group,omitempty"`
	Version       int    `json:"version,omitempty"`
	Updated       string `json:"updated,omitempty"`
	Name          string `json:"name,omitempty"`
	ExternalCode  string `json:"externalCode,omitempty"`
	Archived      bool   `json:"archived,omitempty"`
	PathName      string `json:"pathName,omitempty"`
	Code          string `json:"code,omitempty"`
	ProductFolder struct {
		Meta *Meta `json:"meta,omitempty"`
	} `json:"productFolder,omitempty"`
}

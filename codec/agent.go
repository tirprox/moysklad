package codec

type Agent struct {
	Meta           Meta            `json:"meta,omitempty"`
	ID             string          `json:"id,omitempty"`
	AccountID      string          `json:"accountId,omitempty"`
	Owner          *Owner          `json:"owner,omitempty"`
	Shared         bool            `json:"shared,omitempty"`
	Group          *Group          `json:"group,omitempty"`
	Version        int             `json:"version,omitempty"`
	Updated        string          `json:"updated,omitempty"`
	Name           string          `json:"name,omitempty"`
	ExternalCode   string          `json:"externalCode,omitempty"`
	Archived       bool            `json:"archived,omitempty"`
	Created        string          `json:"created,omitempty"`
	CompanyType    string          `json:"companyType,omitempty"`
	Phone          string          `json:"phone,omitempty"`
	Accounts       *Accounts       `json:"accounts,omitempty"`
	Tags           []string        `json:"tags,omitempty"`
	Contactpersons *Contactpersons `json:"contactpersons,omitempty"`
	Notes          *Notes          `json:"notes,omitempty"`
	State          *State          `json:"state,omitempty"`
	SalesAmount    int             `json:"salesAmount,omitempty"`
}

type Accounts struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Contactpersons struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Notes struct {
	Meta *Meta `json:"meta,omitempty"`
}

type State struct {
	Meta *Meta `json:"meta,omitempty"`
}

type Organization struct {
	Meta *Meta `json:"meta,omitempty"`
}

type OrganizationAccount struct {
	Meta *Meta `json:"meta,omitempty"`
}

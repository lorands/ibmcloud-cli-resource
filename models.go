package resource

import "time"

type Source struct {
	ApiEndpoint   string `json:"api,omitempty"`
	Region        string `json:"region,omitempty"`
	ResourceGroup string `json:"resource_group,omitempty"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AccountId     string `json:"account_id,omitempty"`
	Verbose       bool   `json:"verbose,omitempty"`
}

type Version struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Params struct {
	Cmd     string  `json:"command"`
	SCmd    string  `json:"subcommand"`
	PParams PParams `json:"params,omitempty"`
	Tags    Tags    `json:"tags,omitempty"`
}

type Tags []string

type PParams []string

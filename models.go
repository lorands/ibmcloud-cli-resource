package resource

import "time"

type Source struct {
	Region        string `json:"region"`
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
	Cmd               string  `json:"command"`
	SCmd              string  `json:"subcommand"`
	PParams           PParams `json:"params,omitempty"`
	JSONOutputFileStr string  `json:"jsonOutputFile,omitempty"`
}

type PParams []string

package in

import (
	"github.com/lorands/ibmcloud-cli-resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
	Params  Params           `json:"params"`
}

type Params struct {
	resource.Params
	JSONOutputFileStr string `json:"jsonOutputFile,omitempty"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}

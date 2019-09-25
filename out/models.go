package out

import (
	"github.com/lorands/ibmcloud-cli-resource"
)

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	resource.Params
	Tags Tags `json:"tags,omitempty"`
}

type Tags []string

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}

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
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}

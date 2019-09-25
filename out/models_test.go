package out

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestOutModel(t *testing.T) {
	jsonData := []byte(`
	{
		"params": {
			"command": "resource",
			"subcommand": "group-create",
			"params": [
				"my-resource-group",
				"firstparam",
				"--secondone myfile.yml"
			],
			  "tags": [
				"tag1",
				"tag2"
			  ]
		},
		"source": {
			
		}
	}`)
	var request Request
	dataReader := bytes.NewReader(jsonData)
	if err := json.NewDecoder(dataReader).Decode(&request); err != nil {
		t.Errorf("reading request from stdin: %s", err)
	}

	x:= "resource"
	if request.Params.Cmd != x {
		t.Errorf("Expected: [%s] but got: [%s]", x, request.Params.Cmd )
	}

	x2:= "firstparam"
	if request.Params.PParams[1] != x2 {
		t.Errorf("Expected: [%s] but got: [%s]", x2, request.Params.PParams[1] )
	}

	if len(request.Params.PParams) != 3 {
		t.Errorf("Expected: params size 3 but got: %d", len(request.Params.PParams) )
	}

	x3:= "--secondone myfile.yml"
	if request.Params.PParams[2] != x3 {
		t.Errorf("Expected: [%s] but got: [%s]", x3, request.Params.PParams[2] )
	}

	if len(request.Params.Tags) != 2 {
		t.Errorf("Tags size expected to be 2 but got: [%d]", len(request.Params.Tags))
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestMainRun(t *testing.T) {
	jsonData := []byte(`
	{
		"params": {
			"command": "resource",
			"subcommand": "groups",
			"jsonOutputFile": "output.json"
		},
		"source": {

			"verbose": true
		}
	}`)

	var stdin bytes.Buffer

	stdin.Write(jsonData)

	err := Run("/tmp", &stdin, false)
	if err != nil {
		t.Error(err)
	}
}
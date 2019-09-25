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
    		"tags": [
				"tag1",
				"tag2"
			], 
			"jsonOutputFile": "fokaboka.json"
		},
		"source": {
			"region": "eu-gb", 
			"username": "",
			"password": "",
			"account_id": "",
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

func TestRunFail(t *testing.T){
	jsonData := []byte(`
	{
		"params": {
			"command": "resource",
			"subcommand": "groups",
			"params": [
				"my-resource-group",
				"firstparam"
			],
			  "tags": [
				"tag1",
				"tag2"
			  ], 
			"jsonOutputFile": "fokaboka.json"
		},
		"source": {
			"region": "eu-gb", 
			"username": "",
			"password": "",
			"account_id": "",
			"verbose": true
		}
	}`)

	var stdin bytes.Buffer

	stdin.Write(jsonData)

	err := Run("/tmp", &stdin, false)
	if err == nil {
		t.Error("Test should fail as the args are invalid")
	}
}
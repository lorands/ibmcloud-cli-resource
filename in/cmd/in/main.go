package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	resource "github.com/lorands/ibmcloud-cli-resource"
	"github.com/lorands/ibmcloud-cli-resource/in"
)

var trace bool

func main() {
	if err := Run(os.Args[1], os.Stdin, true); err != nil {
		os.Exit(1)
	}
}

func Run(sourceDir string, stdin io.Reader, doLogin bool) error {
	var request in.Request

	if err := json.NewDecoder(stdin).Decode(&request); err != nil {
		resource.Fatal("reading request from stdin", err)
	}

	trace = request.Source.Verbose

	tracelog("Input directory set: %s.", sourceDir)
	tracelog("Request params set: %v\n", request)

	if doLogin {
		//login
		if err := request.Source.Login(resource.NoTrace); err != nil {
			return err
		}
	}

	pars := resource.CliParams{request.Params.Cmd, request.Params.SCmd}
	for _, p := range request.Params.PParams {
		pars = append(pars, resource.ProcessTemplate(p))
	}

	var workFile string
	if len(request.Params.JSONOutputFileStr) > 0 {
		pars = append(pars, "--output=JSON")
		workFile = sourceDir + "/" + request.Params.JSONOutputFileStr
		tracelog("work file is %s\n", workFile)
	}

	if err := pars.IbmCloudCliRun(workFile, tracelog); err != nil {
		return err
	}

	timestamp := request.Version.Timestamp
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	response := in.Response{
		Version: resource.Version{
			Timestamp: timestamp,
		},
	}
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		resource.Fatal("writing response", err)
	}

	return nil
}

func tracelog(message string, args ...interface{}) {
	if trace {
		_, _ = fmt.Fprintf(os.Stderr, message, args...)
	}
}

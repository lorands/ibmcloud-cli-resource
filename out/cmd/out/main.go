package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	resource "github.com/lorands/ibmcloud-cli-resource"
	"github.com/lorands/ibmcloud-cli-resource/out"
)

var trace bool

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("usage: %v <sources directory>", os.Args[0]))
		os.Exit(1)
	}

	err := Run(os.Args[1], os.Stdin, true)
	if err != nil {
		os.Exit(1)
	}
}

func Run(sourceDir string, stdin io.Reader, doLogin bool) error {

	var request out.Request
	inputRequest(&request, stdin)

	trace = request.Source.Verbose

	tracelog("Input directory set: %s.", sourceDir)
	tracelog("Request params set: %v\n", request)

	if doLogin {
		//login
		if err := request.Source.Login(notrace); err != nil {
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

	metadata := make([]resource.MetadataPair, 1)
	now := time.Now()
	metadata[0] = resource.MetadataPair{
		Name:  "Timestamp",
		Value: now.String(),
	}

	timestamp := time.Now()
	version := resource.Version{
		Timestamp: timestamp,
	}
	//output to stdout...
	response := out.Response{
		Version:  version,
		Metadata: metadata,
	}

	outputResponse(response)

	return nil
}

func inputRequest(request *out.Request, stdin io.Reader) {
	if err := json.NewDecoder(stdin).Decode(request); err != nil {
		log.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response out.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatal("writing response to stdout", err)
	}
}

func notrace(message string, args ...interface{}) {
	//does nothing
}
func tracelog(message string, args ...interface{}) {
	if trace {
		_, _ = fmt.Fprintf(os.Stderr, message, args...)
	}
}

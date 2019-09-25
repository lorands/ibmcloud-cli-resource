package resource

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"text/template"
)

type Trace func(message string, args ...interface{})

type Cli interface {
	Login(tracelog Trace) error
	IbmCloudCliRun(pars []string, workFile string, tracelog Trace) error
}


func (source *Source) Login(tracelog Trace) error {

	region := "eu-gb"
	if len(source.Region) > 0 {
		region = source.Region
	}

	pars := []string {"login",
		"-a cloud.ibm.com",
		fmt.Sprintf("-r %s", region),
		fmt.Sprintf("-u %s", source.Username),
		fmt.Sprintf("-p %s", source.Password),
	}

	if len(source.AccountId)>0 {
		pars = append(pars, source.AccountId)
	}
	if len(source.ResourceGroup) > 0 {
		pars = append(pars, source.ResourceGroup)
	}

	if err := IbmCloudCliRun(pars, "", tracelog); err != nil {
		return err
	}

	return nil
}

func IbmCloudCliRun(pars []string, workFile string, tracelog Trace) error {

	tracelog("About to execute ibmcloud with params: %v", pars)
	cmd := exec.Command("ibmcloud", pars...)
	var sout bytes.Buffer
	cmd.Stdout = &sout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				tracelog("Exit Status: %d", status.ExitStatus())
				if status.ExitStatus() != 0 {
					return fmt.Errorf("Non zero exit code form ibmcloud: %d", status.ExitStatus())
				}
			}
		} else {
			//log.Fatalf("cmd.Wait: %v", err)
			return err
		}
	}
	//do we write out to a file?
	if len(workFile) > 0 {
		if err := ioutil.WriteFile(workFile, sout.Bytes(), 0644); err != nil {
			return fmt.Errorf("fail to write to file: %s", workFile)
		}
	} else {
		tracelog("Output of process is: %s", sout.String())
	}
	return nil
}

//process path from env variables
func ProcessTemplate(tmpl string) string {
	envMap, _ := envToMap()
	t := template.Must(template.New("tmpl").Parse(tmpl))
	var b bytes.Buffer
	_ = t.Execute(&b, envMap)
	return b.String()
}

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	var err error

	for _, v := range os.Environ() {
		split_v := strings.Split(v, "=")
		envMap[split_v[0]] = split_v[1]
	}

	return envMap, err
}

func Fatal(message string, err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}


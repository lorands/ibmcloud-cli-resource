package resource

import (
	"fmt"
	"os"
	"testing"
)

func tracelog(message string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, message, args...)
}

func TestLogin(t *testing.T) {
	source := Source {
		Region: "eu-gb",
		AccountId: "",
		Password: "",
		Username: "",
		Verbose: true,
	}

	err := source.Login(tracelog)
	if err != nil {
		t.Errorf("Fail to run: %v", err)
	}
}

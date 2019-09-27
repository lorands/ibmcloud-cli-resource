package main

import (
	"fmt"
	"os"
	"time"
)

//nothing here
func main() {
	currentVersionTime := time.Time{}
	_, _ = fmt.Fprintf(os.Stdout, "[{\"ref\":\"%s\"}]", currentVersionTime.String())
}

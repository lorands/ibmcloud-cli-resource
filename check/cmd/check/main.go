package main

import (
	"fmt"
	"os"
	"time"
)

//nothing here
func main() {
	currentVersionTime := time.Now()
	_, _ = fmt.Fprintf(os.Stdout, "[{\"ref\":\"%s\"}]", currentVersionTime.String())
}

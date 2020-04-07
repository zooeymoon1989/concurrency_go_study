package main

import (
	flags "github.com/jessevdk/go-flags"
	"os"
)
func main() {

	var opts struct {
		ProjectID string `short:"p" long:"project" description:"Google Cloud Platform Project ID" required:"true"`
		LogName   string `short:"l" long:"logname" description:"The name of the log_server to write to" default:"default"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}



}
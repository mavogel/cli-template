// Package main is the entry point for the CLI application.
package main

import (
	"fmt"
	"os"

	"github.com/mavogel/cli-template/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


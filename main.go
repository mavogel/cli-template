// Package main is the entry point for the CLI application.
package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/mavogel/cli-template/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	if err := fang.Execute(
		context.Background(),
		cmd.RootCmd(),
	); err != nil {
		os.Exit(1)
	}
}


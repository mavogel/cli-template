// Package main is the entry point for the CLI application.
package main

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/charmbracelet/fang"
	"github.com/mavogel/cli-template/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Try to get version info from runtime/debug.BuildInfo first
	if info, ok := debug.ReadBuildInfo(); ok {
		// Get version from module info
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}
		
		// Get commit and date from build settings
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				if setting.Value != "" {
					commit = setting.Value
					if len(commit) > 7 {
						commit = commit[:7] // Use short commit hash
					}
				}
			case "vcs.time":
				if setting.Value != "" {
					date = setting.Value
					// Try to parse and format the date
					if t, err := time.Parse(time.RFC3339, setting.Value); err == nil {
						date = t.Format("2006-01-02")
					}
				}
			}
		}
	}
	
	cmd.SetVersionInfo(version, commit, date)
	if err := fang.Execute(
		context.Background(),
		cmd.RootCmd(),
	); err != nil {
		os.Exit(1)
	}
}


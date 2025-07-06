package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionInfo = struct {
		version string
		commit  string
		date    string
	}{
		version: "dev",
		commit:  "none",
		date:    "unknown",
	}
)

var rootCmd = &cobra.Command{
	Use:   "cli-template",
	Short: "A CLI application template",
	Long:  `A CLI application template built with Go and Cobra for rapid development.`,
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println("Hello from CLI Template!")
		fmt.Println("Use --help to see available commands")
	},
}

// Execute runs the root command and returns any error.
func Execute() error {
	return rootCmd.Execute()
}

// RootCmd returns the root command.
func RootCmd() *cobra.Command {
	return rootCmd
}

// SetVersionInfo sets the version information for the application.
func SetVersionInfo(version, commit, date string) {
	versionInfo.version = version
	versionInfo.commit = commit
	versionInfo.date = date
	rootCmd.Version = fmt.Sprintf("%s (commit: %s, built at: %s)", version, commit, date)
}

func init() {
	// Version flag is now handled by Fang
	// Version info is determined at runtime using runtime/debug.BuildInfo
	// This allows automatic version detection from go.mod and git metadata
}


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
	Run: func(cmd *cobra.Command, args []string) {
		if version, _ := cmd.Flags().GetBool("version"); version {
			fmt.Printf("cli-template %s\n", versionInfo.version)
			fmt.Printf("commit: %s\n", versionInfo.commit)
			fmt.Printf("built at: %s\n", versionInfo.date)
			return
		}
		fmt.Println("Hello from CLI Template!")
		fmt.Println("Use --help to see available commands")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func SetVersionInfo(version, commit, date string) {
	versionInfo.version = version
	versionInfo.commit = commit
	versionInfo.date = date
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}
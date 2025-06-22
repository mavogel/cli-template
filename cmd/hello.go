// Package cmd contains the command-line interface commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const defaultName = "World"

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print a greeting message",
	Long:  `Print a greeting message with optional name parameter.`,
	Run: func(cmd *cobra.Command, _ []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = defaultName
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("name", "n", "", "Name to greet")
}


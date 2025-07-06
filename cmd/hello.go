// Package cmd contains the command-line interface commands.
package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const defaultName = "World"

// HelloAction performs the hello greeting logic.
// It takes a name and a writer, and writes the greeting to the writer.
func HelloAction(name string, w io.Writer) error {
	if name == "" {
		name = defaultName
	}
	_, err := fmt.Fprintf(w, "Hello, %s!\n", name)
	return err
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print a greeting message",
	Long:  `Print a greeting message with optional name parameter.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		name, _ := cmd.Flags().GetString("name")
		return HelloAction(name, cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("name", "n", "", "Name to greet")
}


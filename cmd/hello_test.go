package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestHelloCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "default greeting",
			args:     []string{},
			expected: "Hello, World!",
		},
		{
			name:     "custom name",
			args:     []string{"--name", "Alice"},
			expected: "Hello, Alice!",
		},
		{
			name:     "custom name short flag",
			args:     []string{"-n", "Bob"},
			expected: "Hello, Bob!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			cmd := &cobra.Command{
				Use:   "hello",
				Short: "Print a greeting message",
				Long:  `Print a greeting message with optional name parameter.`,
				Run: func(cmd *cobra.Command, _ []string) {
					name, _ := cmd.Flags().GetString("name")
					if name == "" {
						name = defaultName
					}
					cmd.Printf("Hello, %s!\n", name)
				},
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")

			cmd.SetOut(buf)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if err != nil {
				t.Errorf("Execute() error = %v", err)
			}

			output := buf.String()
			if output != tt.expected+"\n" {
				t.Errorf("Expected output %q, got %q", tt.expected+"\n", output)
			}
		})
	}
}


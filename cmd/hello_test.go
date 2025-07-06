package cmd

import (
	"io"
	"os"
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
			args:     []string{"hello"},
			expected: "Hello, World!\n",
		},
		{
			name:     "custom name",
			args:     []string{"hello", "--name", "Alice"},
			expected: "Hello, Alice!\n",
		},
		{
			name:     "custom name short flag",
			args:     []string{"hello", "-n", "Bob"},
			expected: "Hello, Bob!\n",
		},
		{
			name:     "empty name flag uses default",
			args:     []string{"hello", "--name", ""},
			expected: "Hello, World!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Create a new root command for testing
			rootCmd := &cobra.Command{Use: "test"}
			
			// Add the hello command to the test root
			rootCmd.AddCommand(helloCmd)
			
			// Set args and execute
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			// Restore stdout and read output
			w.Close()
			os.Stdout = oldStdout
			output, _ := io.ReadAll(r)

			if err != nil {
				t.Errorf("Execute() error = %v", err)
			}

			if string(output) != tt.expected {
				t.Errorf("Expected output %q, got %q", tt.expected, string(output))
			}
		})
	}
}

func TestHelloCommandRunFunction(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
		expected string
	}{
		{
			name:     "run with default name",
			flagName: "",
			expected: "Hello, World!\n",
		},
		{
			name:     "run with custom name",
			flagName: "Charlie",
			expected: "Hello, Charlie!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Create a minimal command for testing the Run function directly
			cmd := &cobra.Command{}
			cmd.Flags().StringP("name", "n", tt.flagName, "Name to greet")
			
			// Call the Run function directly
			helloCmd.Run(cmd, []string{})

			// Restore stdout and read output
			w.Close()
			os.Stdout = oldStdout
			output, _ := io.ReadAll(r)

			if string(output) != tt.expected {
				t.Errorf("Expected output %q, got %q", tt.expected, string(output))
			}
		})
	}
}


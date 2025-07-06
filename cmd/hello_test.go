package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

// TestHelloAction tests the HelloAction function directly
func TestHelloAction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "default name",
			input:    "",
			expected: "Hello, World!\n",
			wantErr:  false,
		},
		{
			name:     "custom name",
			input:    "Alice",
			expected: "Hello, Alice!\n",
			wantErr:  false,
		},
		{
			name:     "name with spaces",
			input:    "John Doe",
			expected: "Hello, John Doe!\n",
			wantErr:  false,
		},
		{
			name:     "name with special characters",
			input:    "User@123",
			expected: "Hello, User@123!\n",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			buf := &bytes.Buffer{}
			
			// Call the function
			err := HelloAction(tt.input, buf)
			
			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// Check output
			if got := buf.String(); got != tt.expected {
				t.Errorf("HelloAction() output = %q, want %q", got, tt.expected)
			}
		})
	}
}

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
			if err := w.Close(); err != nil {
				t.Fatalf("Failed to close writer: %v", err)
			}
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

// TestHelloCommandRunE tests the command's RunE function directly
func TestHelloCommandRunE(t *testing.T) {
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
			// Create a buffer to capture output
			buf := &bytes.Buffer{}
			
			// Create a minimal command for testing the RunE function directly
			cmd := &cobra.Command{}
			cmd.SetOut(buf)
			cmd.Flags().StringP("name", "n", tt.flagName, "Name to greet")
			
			// Call the RunE function directly
			err := helloCmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("RunE() error = %v", err)
			}

			if got := buf.String(); got != tt.expected {
				t.Errorf("Expected output %q, got %q", tt.expected, got)
			}
		})
	}
}


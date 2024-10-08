package engine

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

// Scan reads input from TTY if available, otherwise it uses the provided flag.
func Scan(cmd *cobra.Command, args []string, title string) string {
	// Check if TTY is available (i.e., input from terminal)

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		// Prompt the user with the provided title if TTY is available
		if title != "" {
			fmt.Print(title + " : ")
		}

		// Read the user input using bufio
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	// Fallback to using a flag or argument if no TTY is available
	// Check if input flag is provided, using a flag named --input or -i
	inputFlag, _ := cmd.Flags().GetString("input")
	if inputFlag != "" {
		return inputFlag
	}

	// If no input flag provided and TTY is unavailable, return an empty string or handle error
	fmt.Println("No TTY available and no input flag provided.")
	return ""
}

package cmd

import (
	"log"
)

// Execute runs the root command (rootCmd)
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Log the error and exit
		log.Fatalf("Error executing root command: %v", err)
	}
}

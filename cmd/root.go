package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toolbelt",
	Short: "Toolbelt consolidates scripts into a CLI tool",
	Long:  "Toolbelt is a CLI application that consolidates scripts into a single tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toolbelt is a CLI tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

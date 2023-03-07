package cmd

import (
	"os"

	"anddd7.github.com/adcli/cmd/blog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adcli",
	Short: "a colletion of customized cli",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(blog.BlogCmd)
}

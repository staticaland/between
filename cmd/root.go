package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "between",
	Short: "Replace text between markers in files",
	Long: `between is a CLI tool that replaces text between two markers in files.

Examples:
  between replace -f README.md --content "New documentation content"
  between replace -f config.yaml --start-marker "# BEGIN CONFIG" --end-marker "# END CONFIG" --content "key: value"`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
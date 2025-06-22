package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"between/internal/processor"
)

var (
	filename    string
	startMarker string
	endMarker   string
	content     string
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace content between markers in a file",
	Long: `Replace content between two markers in a file.

The command will find the start and end markers in the specified file,
then replace all content between them with the new content.

Example:
  between replace -f README.md --content "New documentation content"
  between replace -f config.yaml --start-marker "# BEGIN" --end-marker "# END" --content "new config"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if filename == "" {
			return fmt.Errorf("filename is required")
		}
		if content == "" {
			return fmt.Errorf("content is required")
		}
		
		return processor.ProcessFile(filename, startMarker, endMarker, content)
	},
}

func init() {
	replaceCmd.Flags().StringVarP(&filename, "file", "f", "", "File to process (required)")
	replaceCmd.Flags().StringVar(&startMarker, "start-marker", "<!-- BEGIN -->", "Start marker")
	replaceCmd.Flags().StringVar(&endMarker, "end-marker", "<!-- END -->", "End marker")
	replaceCmd.Flags().StringVar(&content, "content", "", "New content to insert (required)")
	
	replaceCmd.MarkFlagRequired("file")
	replaceCmd.MarkFlagRequired("content")
}
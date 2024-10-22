package cmd

import (
    "fmt"
		"os"
    "github.com/atotto/clipboard"
    "github.com/spf13/cobra"
    "cue/internal"
)

func init() {
    contentCmd.Flags().BoolP("strip-tags", "s", false, "Strip XML tags from the output")
    contentCmd.Flags().BoolP("copy", "c", false, "Copy the output to clipboard instead of printing")
    rootCmd.AddCommand(contentCmd)
}

var contentCmd = &cobra.Command{
    Use:   "content [prompt name]",
    Short: "Get the content of a prompt",
    Long: `Get the content of a prompt. By default, shows content with internal XML tags.
Use -s or --strip-tags to remove all XML tags from the output.
Use -c or --copy to copy the content to your clipboard instead of printing to stdout.`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
			stripTags, _ := cmd.Flags().GetBool("strip-tags")
			copyToClipboard, _ := cmd.Flags().GetBool("copy")
			promptName := args[0]

			prompt, err := internal.LoadPrompt(promptName + ".xml")
			if err != nil {
					return fmt.Errorf("failed to load prompt '%s': %w", promptName, err)
			}

			content := prompt.Content
			if stripTags {
					content = internal.StripXMLTags(content)
			}

			if copyToClipboard {
					if err := clipboard.WriteAll(content); err != nil {
							return fmt.Errorf("failed to copy to clipboard: %w", err)
					}
					fmt.Fprintln(os.Stderr, "Content copied to clipboard")
			} else {
					fmt.Println(content)
			}

			return nil
	},
}

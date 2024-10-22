package cmd

import (
    "fmt"
    "cue/internal"
    "github.com/spf13/cobra"
)

func init() {
    contentCmd.Flags().BoolP("strip-tags", "s", false, "Strip XML tags from the output")
    rootCmd.AddCommand(contentCmd)
}

var contentCmd = &cobra.Command{
    Use:   "content [prompt name]",
    Short: "Get the content of a prompt",
    Long: `Get the content of a prompt. By default, shows content with internal XML tags.
Use -s or --strip-tags to remove all XML tags from the output.`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
			stripTags, _ := cmd.Flags().GetBool("strip-tags")
			promptName := args[0]

			prompt, err := internal.LoadPrompt(promptName + ".xml")
			if err != nil {
					return fmt.Errorf("failed to load prompt '%s': %w", promptName, err)
			}

			if stripTags {
					fmt.Println(internal.StripXMLTags(prompt.Content))
			} else {
					fmt.Println(prompt.Content)
			}

			return nil
	},
}

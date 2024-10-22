package cmd

import (
    "fmt"
		"path/filepath"
    "cue/internal"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new prompt",
	Long:  `Add a new prompt to the collection of AI prompts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
			prompt, err := internal.GetFromUser()
			if err != nil {
					return fmt.Errorf("failed to get prompt information: %w", err)
			}

			if err := prompt.Save(); err != nil {
					return fmt.Errorf("failed to save prompt: %w", err)
			}

			fmt.Printf("Prompt saved successfully to %s\n", filepath.Join(internal.DefaultPromptsDir, prompt.Filename()))
			return nil
	},
}

package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "cue",
    Short: "cue is a CLI tool for managing AI prompts",
    Long: `cue allows you to store, retrieve, and manage AI prompts from your terminal.
    You can add new prompts, view existing ones, list all saved prompts, and get the content of a prompt.`,
}

func init() {
	// Disable completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Disable the help command (but keep the -h flag)
	rootCmd.SetHelpCommand(&cobra.Command{
			Use:    "no-help",
			Hidden: true,
	})
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

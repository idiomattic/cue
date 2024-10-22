package cmd

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "cue",
    Short: "cue is a CLI tool for managing AI prompts",
    Long: `cue allows you to store, retrieve, and manage AI prompts from your terminal.
    You can add new prompts, view existing ones, list all saved prompts, and get the content of a prompt.`,
}

// DefaultPromptsDir is the default directory for storing prompts
var DefaultPromptsDir = filepath.Join(os.Getenv("HOME"), "prompt_engineer", "prompt_files")

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

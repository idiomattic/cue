package main

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "cue",
    Short: "cue is a CLI tool for managing AI prompts",
    Long: `cue allows you to store, retrieve, and manage AI prompts from your terminal.
    You can add new prompts, view existing ones, list all saved prompts, and get the content of a prompt.`,
}

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new prompt",
    Long:  `Add a new prompt to the collection of AI prompts.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Adding a new prompt...")
        // We'll implement the actual functionality here later
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "cue/internal"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all available prompts",
    Long:  `List all prompts stored in your prompts directory.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        files, err := os.ReadDir(internal.DefaultPromptsDir)
        if err != nil {
            if os.IsNotExist(err) {
                fmt.Println("No prompts found. Use 'cue add' to create your first prompt.")
                return nil
            }
            return fmt.Errorf("failed to read prompts directory: %w", err)
        }

        if len(files) == 0 {
            fmt.Println("No prompts found. Use 'cue add' to create your first prompt.")
            return nil
        }

        fmt.Println("Available prompts:")
        for _, file := range files {
            if filepath.Ext(file.Name()) != ".xml" {
                continue
            }

            prompt, err := internal.LoadPrompt(file.Name())
            if err != nil {
                fmt.Printf("Warning: Could not read prompt %s: %v\n", file.Name(), err)
                continue
            }

            fmt.Printf("%s: %s (Category: %s)\n",
                file.Name()[:len(file.Name())-4], // remove .xml
                prompt.Title,
                prompt.Category,
            )
        }

        return nil
    },
}

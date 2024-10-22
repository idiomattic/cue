package internal

import (
    "os"
    "path/filepath"
    "strings"
    "time"
)

var DefaultPromptsDir = filepath.Join(os.Getenv("HOME"), ".cues")

// Prompt represents a stored prompt
type Prompt struct {
    Title        string
    Content      string
    Category     string
    Tags         []string
    CreatedDate  time.Time
    LastModified time.Time
}

// Filename generates a filename for the prompt
func (p *Prompt) Filename() string {
    // Convert title to lowercase and replace spaces with underscores
    sanitized := strings.Map(func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return r + 32 // convert to lowercase
        case r >= 'a' && r <= 'z':
            return r
        case r >= '0' && r <= '9':
            return r
        default:
            return '_'
        }
    }, p.Title)
    return sanitized + ".xml"
}

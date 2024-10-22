package internal

import (
	"time"
    "encoding/xml"
    "fmt"
		"strings"
    "os"
    "path/filepath"
)

// xmlPrompt is used for XML marshaling/unmarshaling
type xmlPrompt struct {
    XMLName      xml.Name `xml:"prompt"`
    Title        string   `xml:"title"`
    Content      string   `xml:"content"`
    Category     string   `xml:"category"`
    Tags         []string `xml:"tags>tag"`
    CreatedDate  string   `xml:"created_date"`
    LastModified string   `xml:"last_modified"`
}

// Save writes the prompt to a file in XML format
func (p *Prompt) Save() error {
    // Ensure the prompts directory exists
    if err := os.MkdirAll(DefaultPromptsDir, 0755); err != nil {
        return fmt.Errorf("failed to create prompts directory: %w", err)
    }

    // Convert to XML format
    xp := xmlPrompt{
        Title:        p.Title,
        Content:      p.Content,
        Category:     p.Category,
        Tags:         p.Tags,
        CreatedDate:  p.CreatedDate.Format("2006-01-02"),
        LastModified: p.LastModified.Format("2006-01-02"),
    }

    // Marshal to XML
    data, err := xml.MarshalIndent(xp, "", "    ")
    if err != nil {
        return fmt.Errorf("failed to marshal prompt to XML: %w", err)
    }

    // Add XML header
    xmlData := []byte(xml.Header + string(data))

    // Write to file
    filename := filepath.Join(DefaultPromptsDir, p.Filename())
    if err := os.WriteFile(filename, xmlData, 0644); err != nil {
        return fmt.Errorf("failed to write prompt file: %w", err)
    }

    return nil
}

// LoadPrompt reads a prompt from an XML file
func LoadPrompt(filename string) (*Prompt, error) {
	data, err := os.ReadFile(filepath.Join(DefaultPromptsDir, filename))
	if err != nil {
			return nil, fmt.Errorf("failed to read prompt file: %w", err)
	}

	var xp xmlPrompt
	if err := xml.Unmarshal(data, &xp); err != nil {
			return nil, fmt.Errorf("failed to parse prompt XML: %w", err)
	}

	createdDate, err := time.Parse("2006-01-02", xp.CreatedDate)
	if err != nil {
			return nil, fmt.Errorf("failed to parse created date: %w", err)
	}

	lastModified, err := time.Parse("2006-01-02", xp.LastModified)
	if err != nil {
			return nil, fmt.Errorf("failed to parse last modified date: %w", err)
	}

	return &Prompt{
			Title:        xp.Title,
			Content:      xp.Content,
			Category:     xp.Category,
			Tags:         xp.Tags,
			CreatedDate:  createdDate,
			LastModified: lastModified,
	}, nil
}

// StripXMLTags removes XML tags from content while preserving the text
func StripXMLTags(content string) string {
	var inTag bool
	var result strings.Builder

	for i := 0; i < len(content); i++ {
			switch content[i] {
			case '<':
					inTag = true
			case '>':
					inTag = false
			default:
					if !inTag {
							result.WriteByte(content[i])
					}
			}
	}

	return result.String()
}

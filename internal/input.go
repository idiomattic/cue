package internal

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

// GetFromUser prompts the user for prompt information
func GetFromUser() (*Prompt, error) {
    reader := bufio.NewReader(os.Stdin)

    // Get title
    fmt.Print("Enter prompt title: ")
    title, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    title = strings.TrimSpace(title)

    // Get content
    fmt.Println("Enter prompt content (press Ctrl+D when finished):")
    var contentBuilder strings.Builder
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        contentBuilder.WriteString(line)
    }
    content := strings.TrimSpace(contentBuilder.String())

    // Get category (need a new reader since we used Ctrl+D)
    reader = bufio.NewReader(os.Stdin)
    fmt.Print("Enter category (default: General): ")
    category, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    category = strings.TrimSpace(category)
    if category == "" {
        category = "General"
    }

    // Get tags
    fmt.Print("Enter tags (comma-separated): ")
    tagsInput, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    tagsInput = strings.TrimSpace(tagsInput)
    tags := strings.Split(tagsInput, ",")
    for i := range tags {
        tags[i] = strings.TrimSpace(tags[i])
    }

    return &Prompt{
        Title:        title,
        Content:      content,
        Category:     category,
        Tags:         tags,
        CreatedDate:  time.Now(),
        LastModified: time.Now(),
    }, nil
}

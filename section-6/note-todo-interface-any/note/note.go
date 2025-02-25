package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// New returns a new *Note or error if any
func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input")
	}
	return &Note{Title: title, Content: content, CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
}

// Display displays note's details
func (n *Note) Display() {
	fmt.Println(n.Title)
	fmt.Println(n.Content)
	fmt.Println(n.CreatedAt)
	fmt.Println(n.UpdatedAt)
}

// Save save note to a file
func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// String implements Stringer interface
func (n *Note) String() string {
	return fmt.Sprintf("Title: %s, Content: %s, CreatedAt: %s, UpdatedAt: %s", n.Title, n.Content, n.CreatedAt, n.UpdatedAt)
}

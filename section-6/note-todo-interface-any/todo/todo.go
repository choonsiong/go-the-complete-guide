package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

const todoFilename = "todo.json"

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid input")
	}
	return &Todo{text}, nil
}

func (t *Todo) Display() {
	fmt.Println(t.Text)
}

func (t *Todo) Save() error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFilename, data, 0644)
}

func (t *Todo) String() string {
	return fmt.Sprintf("Text: %s", t.Text)
}

package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// New returns a new *FileManager
func New(input, output string) *FileManager {
	return &FileManager{
		InputFilePath:  input,
		OutputFilePath: output,
	}
}

// ReadLines reads lines from the given file and returns a slice of string
// or error if any
func (m *FileManager) ReadLines() ([]string, error) {
	var lines []string

	f, err := os.Open(m.InputFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// WriteResult is a wrapper method for WriteJson
func (m *FileManager) WriteResult(data any) error {
	return m.WriteJson(data)
}

// WriteJson writes data to the file path specifies in FileManager.OutputFilePath
func (m *FileManager) WriteJson(data any) error {
	f, err := os.Create(m.OutputFilePath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(f)
	return encoder.Encode(data)
}

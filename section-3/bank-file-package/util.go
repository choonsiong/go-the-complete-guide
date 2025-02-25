package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readFloatFromFile(filename string) (float64, error) {
	byteArr, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If file not found, create a new one with 0 value
			return 0, os.WriteFile(filename, []byte(fmt.Sprint("0")), 0644)
		} else {
			return 0, err
		}
	}

	return strconv.ParseFloat(string(byteArr), 64)
}

func writeFloatToFile(value float64, filename string) error {
	return os.WriteFile(filename, []byte(fmt.Sprint(value)), 0644)
}

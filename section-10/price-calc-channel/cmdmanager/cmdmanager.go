package cmdmanager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdManager struct {
}

func New() *CmdManager {
	return &CmdManager{}
}

// ReadLines reads from stdin and returns a []string and error if any
func (c *CmdManager) ReadLines() ([]string, error) {
	fmt.Print("Enter prices (separated by whitespace): ")

	var prices []string

	floats, err := getFloats()
	if err != nil {
		return nil, err
	}

	for _, f := range floats {
		prices = append(prices, fmt.Sprintf("%.2f", f))
	}

	return prices, nil
}

// WriteResult writes data to stdout
func (c *CmdManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

// getFloats read a single line from stdin and returns []float64 or error
// if any
func getFloats() ([]float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimSpace(input)
	ss := strings.Split(input, " ")

	var floats []float64

	for _, s := range ss {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}

	return floats, nil
}

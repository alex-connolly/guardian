package util

import "fmt"

// Error ...
type Error struct {
	LineNumber int
	Message    string
}

// Errors ...
type Errors []Error

// Format ...
func (e Errors) Format() string {
	whole := ""
	whole += fmt.Sprintf("%d errors\n", len(e))
	for _, err := range e {
		whole += fmt.Sprintf("Line %d: %s\n", err.LineNumber, err.Message)
	}
	return whole
}

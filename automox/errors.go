package automox

import (
	"errors"
	"fmt"
)

// ErrorResponse represents a Automox API error
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Helper to be used for API client config errors
func missingClientConfigErr(attr string) error {
	errTxt := fmt.Sprintf("A valid Automox %s is required to create a new API client", attr)
	return errors.New(errTxt)
}

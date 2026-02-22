package check

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrMissingStatus = errors.New("missing status")
	ErrInvalidStatus = errors.New("invalid status")
)

// ValidateSmokeOutput ensures the top-level status field exists and is valid.
func ValidateSmokeOutput(payload []byte) error {
	var obj map[string]any
	if err := json.Unmarshal(payload, &obj); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	status, ok := obj["status"]
	if !ok {
		return ErrMissingStatus
	}

	s, ok := status.(string)
	if !ok || s == "" {
		return ErrInvalidStatus
	}
	if s != "pass" && s != "fail" {
		return ErrInvalidStatus
	}

	return nil
}

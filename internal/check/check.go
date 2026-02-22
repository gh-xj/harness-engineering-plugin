package check

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrMissingStatus = errors.New("missing status")
	ErrInvalidStatus = errors.New("invalid status")
	ErrMissingRunID = errors.New("missing run_id")
	ErrInvalidRunID = errors.New("invalid run_id")
	ErrMissingChecks = errors.New("missing checks")
	ErrInvalidChecks = errors.New("invalid checks")
)

// ValidateSmokeOutput ensures required top-level fields exist and are valid.
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

	runID, ok := obj["run_id"]
	if !ok {
		return ErrMissingRunID
	}

	r, ok := runID.(string)
	if !ok || r == "" {
		return ErrInvalidRunID
	}

	checks, ok := obj["checks"]
	if !ok {
		return ErrMissingChecks
	}
	if _, ok := checks.([]any); !ok {
		return ErrInvalidChecks
	}

	return nil
}

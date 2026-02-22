package check

import (
	"errors"
	"testing"
)

func TestValidateSmokeOutput_ValidPayload(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"status":"pass","run_id":"run-1","checks":[]}`)
	if err := ValidateSmokeOutput(payload); err != nil {
		t.Fatalf("expected valid payload, got error: %v", err)
	}
}

func TestValidateSmokeOutput_MissingStatus(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"run_id":"run-1","checks":[]}`)
	if err := ValidateSmokeOutput(payload); err == nil {
		t.Fatal("expected error when status is missing")
	}
}

func TestValidateSmokeOutput_InvalidStatusEnum(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"status":"unknown","run_id":"run-1","checks":[]}`)
	err := ValidateSmokeOutput(payload)
	if !errors.Is(err, ErrInvalidStatus) {
		t.Fatalf("expected ErrInvalidStatus, got: %v", err)
	}
}

func TestValidateSmokeOutput_MissingRunID(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"status":"pass","checks":[]}`)
	err := ValidateSmokeOutput(payload)
	if !errors.Is(err, ErrMissingRunID) {
		t.Fatalf("expected ErrMissingRunID, got: %v", err)
	}
}

func TestValidateSmokeOutput_MissingChecks(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"status":"pass","run_id":"run-1"}`)
	err := ValidateSmokeOutput(payload)
	if !errors.Is(err, ErrMissingChecks) {
		t.Fatalf("expected ErrMissingChecks, got: %v", err)
	}
}

func TestValidateSmokeOutput_MalformedJSON(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"status":"pass"`)
	if err := ValidateSmokeOutput(payload); err == nil {
		t.Fatal("expected error for malformed JSON")
	}
}

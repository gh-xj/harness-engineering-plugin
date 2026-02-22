package check

import "testing"

func TestValidateSmokeOutput_MissingStatus(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"run_id":"run-1","checks":[]}`)
	if err := ValidateSmokeOutput(payload); err == nil {
		t.Fatal("expected error when status is missing")
	}
}

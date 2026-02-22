package main

import (
	"fmt"
	"io"
	"os"

	"../../internal/check"
)

func main() {
	payload, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read stdin:", err)
		os.Exit(1)
	}

	if err := check.ValidateSmokeOutput(payload); err != nil {
		fmt.Fprintln(os.Stderr, "validation error:", err)
		os.Exit(1)
	}
}

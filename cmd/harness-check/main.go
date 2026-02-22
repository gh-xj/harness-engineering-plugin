package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

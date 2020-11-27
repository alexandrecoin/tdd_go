package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	result := buffer.String()
	expectedResult := "Hello, Chris"

	if result != expectedResult {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}

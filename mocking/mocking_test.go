package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)

	result := buffer.String()
	expectedResult := "3"

	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}

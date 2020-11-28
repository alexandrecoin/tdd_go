package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	result := buffer.String()
	expectedResult := `3
2
1
Go!`

	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("Expected 3 sleep calls but got %d", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

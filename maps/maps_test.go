package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	result := Search(dictionary, "test")
	expectedResult := "this is just a test"
	assertStrings(t, result, expectedResult)
}

func assertStrings(t *testing.T, result, expectedResult string) {
	t.Helper()
	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}
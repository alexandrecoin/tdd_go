package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	result := dictionary.Search("test")
	expectedResult := "this is just a test"
	assertStrings(t, result, expectedResult)
}

func assertStrings(t *testing.T, result, expectedResult string) {
	t.Helper()
	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}
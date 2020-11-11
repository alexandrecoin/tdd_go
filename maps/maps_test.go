package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		t.Parallel()
		dictionary := Dictionary{"test": "this is just a test"}
	
		result, _ := dictionary.Search("test")
		expectedResult := "this is just a test"
		assertStrings(t, result, expectedResult)
	})

	t.Run("unknown word", func(t *testing.T) {
		t.Parallel()
		dictionary := Dictionary{"test": "this is another test"}
		_, err := dictionary.Search("someWord")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t *testing.T, result, expectedResult string) {
	t.Helper()
	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}

func assertError(t *testing.T, result, expectedResult error) {
	t.Helper()
	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}
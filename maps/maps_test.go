package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
	
		result, _ := dictionary.Search("test")
		expectedResult := "this is just a test"
		assertStrings(t, result, expectedResult)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is another test"}
		_, err := dictionary.Search("someWord")
		expectedResult := "Could not find the word you are looking for."

		if err == nil {
			t.Fatal("Expected to get an error.")
		}

		assertStrings(t, err.Error(), expectedResult)
	})
}

func assertStrings(t *testing.T, result, expectedResult string) {
	t.Helper()
	if expectedResult != result {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}
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

func TestAddEntry(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "entryTest"
		definition := "This is a test for an entry add"
		dictionary.AddEntry(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Already existing word", func(t *testing.T) {
		word := "test"
		definition := "This is a definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.AddEntry(word, "New definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdateEntry(t *testing.T) {
	t.Run("Already existing entry", func (t *testing.T) {
		t.Parallel()
		entry := "test"
		definition := "This is just a test"
		newDefinition := "This is a new definition"
		dictionary := Dictionary{entry: definition}
		err := dictionary.UpdateEntry(entry, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, entry, newDefinition)
	})

	t.Run("New entry", func (t *testing.T) {
		t.Parallel()
		entry := "word"
		definition := "This is just a test"
		dictionary := Dictionary{}
		err :=	dictionary.UpdateEntry(entry, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t *testing.T, result, expectedResult string) {
	t.Helper()
	if result != expectedResult {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}
}

func assertError(t *testing.T, result, expectedResult error) {
	t.Helper()

	if result != expectedResult {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}

	if result == nil {
		if expectedResult == nil {
			return
		}
		t.Fatal("Expected to get an error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word: ", err)
	}

	if definition != result {
		t.Errorf("Expected %q but got %q", definition, result)
	}
}

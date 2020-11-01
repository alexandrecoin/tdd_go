package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when no name is supplied", func (t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish, it says 'Hola, + name'", func (t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French, it says 'Bonjour, + name", func(t *testing.T) {
		got := Hello("Alexandre", "French")
		want := "Bonjour, Alexandre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese, it says 'Ohayou, + name", func (t *testing.T) {
		got := Hello("Tanaka", "Japanese")
		want := "Ohayou, Tanaka"
		assertCorrectMessage(t, got, want)
	})
}

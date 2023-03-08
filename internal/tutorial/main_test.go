package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anddd7", ENGLISH)
		want := "Hello, Anddd7!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("", ENGLISH)
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", SPANISH)
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Anddd7", FRENCH)
		want := "Bonjour, Anddd7!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tell the test suite that this method is a helper.
	// the line number reported will be in our function call rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

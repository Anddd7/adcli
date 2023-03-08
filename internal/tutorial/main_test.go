package main

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anddd7")
		want := "Hello, Anddd7!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

package main

import "testing"

func Test_Hello(t *testing.T) {
	got := Hello("Anddd7")
	want := "Hello, Anddd7!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

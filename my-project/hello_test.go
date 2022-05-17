package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Charles", "")
		want := "Hello, Charles"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Solene", "Spanish")
		want := "Hola, Solene"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sherlock", "French")
		want := "Bonjour, Sherlock"
		assertCorrectMessage(t, got, want)
	})
}

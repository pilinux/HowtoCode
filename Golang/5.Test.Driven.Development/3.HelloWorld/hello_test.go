package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper() // Inform the test-suite that assertCorrectMessage is a helper method
		if got != want {
			t.Errorf("got -->'%s'<--, want -->'%s'<--", got, want)
		}
	}

	t.Run("Greet the person in English", func(t *testing.T) {
		got := Hello("Jacinda", "English")
		want := "Hello, Jacinda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the person in German", func(t *testing.T) {
		got := Hello("Jacinda", "German")
		want := "Hallo, Jacinda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the person in French", func(t *testing.T) {
		got := Hello("Jacinda", "French")
		want := "Bonjour, Jacinda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the person in Spanish", func(t *testing.T) {
		got := Hello("Jacinda", "Spanish")
		want := "Hola, Jacinda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the person in the default language", func(t *testing.T) {
		got := Hello("Jacinda", "")
		want := "Hello, Jacinda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the whole world when empty strings are given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the world in English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the world in German", func(t *testing.T) {
		got := Hello("", "German")
		want := "Hallo, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the world in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet the world in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})
}

package main

import "testing"

func TestHello (t *testing.T) {
	t.Run("Saying hello to various people", func(t *testing.T) {	
		got := Hello("Alejandro", "")
		want := "Hello, Alejandro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Pim", "Spanish")
		want := "Hola, Pim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Charlie", "French")
		want := "Bonjour, Charlie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

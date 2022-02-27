package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("greeting people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
	
		assertCorrectMessage(t, got, want)
	})

	t.Run("spannish greeting", func(t *testing.T) {
		got := Hello("Julia", "SP")
		want := "Hola, Julia"

		assertCorrectMessage(t, got, want)
	})

	t.Run("french greeting", func(t *testing.T) {
		got := Hello("Michel", "FR")
		want := "Bonjour, Michel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("no argument passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
	
		assertCorrectMessage(t, got, want)
	})
}
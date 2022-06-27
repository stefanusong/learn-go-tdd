package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		name := "Bambang"
		got := Hello(name, "")
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in other languages", func(t *testing.T) {
		for key, value := range Prefixes {
			got := Hello("Bambang", key)
			want := value + ", Bambang"
			assertCorrectMessage(t, got, want)
		}
	})

}

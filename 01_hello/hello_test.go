package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello Prs"
		got := Hello("Prs", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		want := "Hello World"
		got := Hello("", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("in Tamil", func(t *testing.T) {
		want := "Vaadaa Maapla"
		got := Hello("Maapla", "Tamil")
		assertCorrectMessage(t, want, got)
	})

	t.Run("in Hindi", func(t *testing.T) {
		want := "Aayiye Maapla"
		got := Hello("Maapla", "Hindi")
		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got: %q", want, got)
	}
}

func ExampleHello() {
	greeting := Hello("Prs", "")
	fmt.Println(greeting)
	// Output: Hello Prs
}

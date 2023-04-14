package main

import (
	"testing"
)

// it's good to use TB interface, used by *testing.T and *testing.B so we can call helper funcs from a test
func assertCorrectMessage(t testing.TB, got, want string) {
	// t Helper marks this function as a helper, if test fails it will mark the line inside the actual test func and not this helper func, so it's easier to debug
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello with a name", func(t *testing.T){
	got := sayHello("Kailey", "")
	want := "Hello Kailey"

	assertCorrectMessage(t, got, want)
})
	t.Run("say Hello World if empty string provided", func(t *testing.T) {
	got := sayHello("", "")
	want := "Hello World"

	assertCorrectMessage(t, got, want)

})

t.Run("in Spanish", func(t *testing.T){
	got := sayHello("Kailey", "Spanish")
	want := "Hola Kailey"

	assertCorrectMessage(t, got, want)

})
}


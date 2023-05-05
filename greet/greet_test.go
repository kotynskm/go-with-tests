package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	// Buffer type from bytes pkg implements Writer interface, and has method Write
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
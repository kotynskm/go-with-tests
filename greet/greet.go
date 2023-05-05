package greet

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	// Printf defaults to stdout
	// fmt.Printf("Hello, %s", name)

	// use writer to send greeting to the buffer
	fmt.Fprintf(writer, "Hello, %s", name)
}
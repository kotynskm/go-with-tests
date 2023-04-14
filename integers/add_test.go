package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("should sum two numbers", func(t *testing.T)  {
		got := Add(2,2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	
}

// example, these are executed just like tests and reflect what the function should actually do
// examples will not run without the Output comment!
func ExampleAdd() {
	sum := Add(2,2)
	fmt.Println(sum)
	// Output: 4
}


package iteration

import (
	"fmt"
	"testing"
)
func TestIterateFive(t *testing.T){
	got := RepeatChar("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

// benchmarks, run like tests and the go framework determines values for us
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChar("a", 5)
	}
}

func ExampleRepeatChar(){
	repeat := RepeatChar("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
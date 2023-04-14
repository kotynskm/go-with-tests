package sum

import (
	"reflect"
	"testing"
)


func TestSum(t *testing.T){

	t.Run("sum an array of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("sum is %d expected %d, with numbers array %v", sum, expected, numbers)
		}
	})
	}

func TestSumAll(t *testing.T){
	// defining a func as variable shown here, since it's declared inside the test it can only be used here
	checkSums := func(t testing.TB, sum, expected []int) {
		t.Helper()
	
		// use with slices since we can't use equality operators
		if !reflect.DeepEqual(sum, expected){
			t.Errorf("sum is %v expected %v", sum, expected)
		}
	
	}
	t.Run("sum all numbers from each slice and return sums in a new slice", func(t *testing.T) {
		numbers := []int{1,2,3}
		numbers2 := []int{2,3}

		sum := SumAllTails(numbers, numbers2)
		expected := []int{5,3}

		// use with slices, we can't use equality operator with slices. note - reflect is not type safe
		checkSums(t, sum, expected)
	})

	t.Run("sum an empty slice", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{2,3})
		expected := []int{0,3}

		checkSums(t, sum, expected)
	})

}

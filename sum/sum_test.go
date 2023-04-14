package sum

import (
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

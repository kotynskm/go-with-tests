package geometry

import (
	"testing"
)


func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0,10.0}
	got := calcPerimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

// func TestArea(t *testing.T) {
// 	t.Run("calc rectangle area", func (t *testing.T)  {
// 	rectangle := Rectangle{10.0,10.0}
// 	got := rectangle.CalcArea()
// 	expected := 100.0

// 	if got != expected {
// 		t.Errorf("got %.2f expected %.2f", got, expected)
// 	}

// 	t.Run("calc circle area", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		got := circle.CalcArea()
// 		expected := 314.1592653589793

// 		if got != expected {
// 			t.Errorf("got %g expected %g", got, expected)
// 		}
// 	})
// 	})
	
// }


// refactored test that uses interface
// func TestArea(t *testing.T) {

// 	// declare helper func inside the test, we want to be able to pass shapes to it, will fail if try to pass anything that is not a shape
// 	checkArea := func(t testing.TB, shape Shape, want float64){
// 		// t Helper marks this function as a helper, if test fails it will mark the line inside the actual test func and not this helper func, so it's easier to debug
// 		t.Helper()
// 		got := shape.CalcArea()
		
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("calc area for rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{10.0,10.0}
// 		checkArea(t, rectangle, 100.0)
// 	})

// 	t.Run("calc area for circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// table driven test example
// run specific tests with a command like go test -run TestArea/Rectangle
func TestArea(t *testing.T){
	areaTests := []struct {
		name string
		shape Shape
		want float64
	} {
		// fields can be optionally named as such, to make the code more clear
		{name: "Rectangle",shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12,6}, 36.0},
	}

	for _, tt := range areaTests {
		// use tt.name to use it as a t.Run test name
		t.Run(tt.name, func (t *testing.T)  {
			got := tt.shape.CalcArea()

		// %#v will print our struct and values so it's easier to see the properties being tested
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
		})
	}
}
package interfaces

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) { // checks a shape is passed,shape interface is passed as an argument
		t.Helper()          // read on what a helper function is
		got := shape.Area() //calls interface , Area() method defines the criteria for something to be a shape
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

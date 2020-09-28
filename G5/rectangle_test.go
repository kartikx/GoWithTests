package rectangle

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}

	got := rectangle.Perimeter()
	expected := 40.0

	if got != expected {
		t.Errorf("Expected %.2f, Got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {

	checkArea := func (t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("Expected %g, Got %g", expected, got)
		}
	}

	t.Run("Testing Rectangle Area: ", func (t *testing.T) {
		rectangle := Rectangle{width: 10.0, height: 10.0}
		
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Testing Circle Area: ", func (t *testing.T) {
		circle := Circle{radius: 10.0}
		
		checkArea(t, circle, 314.1592653589793) 
	})
}
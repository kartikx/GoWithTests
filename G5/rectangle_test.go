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

	areaTests := []struct {
		shape Shape
		expected float64
	}{
		{shape: Rectangle{width: 10.0, height: 10.0 }, expected: 100.0},
		{shape: Circle{radius: 10.0}, expected: 314.1592653589793},
	}

	for _, tt := range(areaTests) {
		got := tt.shape.Area()

		if got != tt.expected {
			t.Errorf("Expected %g, Got %g", tt.expected, got)
		}
	}
}
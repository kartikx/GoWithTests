package shapes

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
		name string
		shape Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 10.0, height: 10.0 }, expected: 100.0},
		{name: "Circle", shape: Circle{radius: 10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12.0, height: 24.0}, expected: 144.0},
	}

	for _, tt := range(areaTests) {
		t.Run(tt.name, func (t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("Expected %g, Got %g, %T", tt.expected, got, tt.shape)
			}
		})
	}
}
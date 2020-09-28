package shapes

import ("math")

// Shape interface.
type Shape interface {
	Area() float64
}

// Rectangle struct.
type Rectangle struct {
	width, height float64
}

// Perimeter returns the Perimeter of the Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Area returns the Area of the Rectangle.
func (r Rectangle) Area() float64 {
	return (r.width * r.height)
}

// Circle struct.
type Circle struct {
	radius float64
}

// Area for Circle.
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// Triangle struct.
type Triangle struct {
	height, base float64
}

// Area for Triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
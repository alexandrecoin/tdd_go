package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 12.0}
	result := Perimeter(rectangle)
	expectedResult := 44.0
	if expectedResult != result {
		t.Errorf("Expected %.2f but got %.2f", expectedResult, result)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name           string
		shape          Shape
		expectedResult float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedResult: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedResult: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedResult: 36.0},
	}

	for _, item := range areaTests {
		t.Run(item.name, func(t *testing.T) {
			result := item.shape.Area()
			if item.expectedResult != result {
				t.Errorf("%#v - Expected %g but got %g result", item.shape, item.expectedResult, result)
			}
		})
	}
}

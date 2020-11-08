package main

import "testing"

func TestPerimeter(t * testing.T) {
	rectangle := Rectangle{10.0, 12.0}
	result := Perimeter(rectangle)
	expectedResult := 44.0
	if expectedResult != result {
		t.Errorf("Expected %.2f but got %.2f", expectedResult, result)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		expectedResult float64
	} {
		{ Rectangle {12, 6}, 72.0 },
		{ Circle {10}, 314.1592653589793 },
	}

	for _, item := range areaTests {
		result := item.shape.Area()
		if item.expectedResult != result {
			t.Errorf("Expected %g but got %g result", item.expectedResult, result)
		}
	}
}
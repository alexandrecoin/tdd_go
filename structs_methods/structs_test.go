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

	checkArea := func(t *testing.T, shape Shape, expectedResult float64) {
		t.Helper()
		result := shape.Area()
		if result != expectedResult{
			t.Errorf("Expected %g but got %g", expectedResult, result)
		}
	} 

	t.Run("Calculate area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 12.0}
		expectedResult := 120.0
		checkArea(t, rectangle, expectedResult)
	})

	t.Run("Calculate the area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		expectedResult := 314.1592653589793
		checkArea(t, circle, expectedResult)
	})
}
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
	rectangle := Rectangle{10.0, 12.0}
	result := Area(rectangle)
	expectedResult := 120.0
	if expectedResult != result {
		t.Errorf("Expected %.2f but got %.2f", expectedResult, result)
	}
}
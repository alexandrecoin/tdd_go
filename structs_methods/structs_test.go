package main

import "testing"

func TestPerimeter(t * testing.T) {
	result := Perimeter(4, 2)
	expectedResult := 12.0
	if expectedResult != result {
		t.Errorf("Expected %.2f but got %.2f", expectedResult, result)
	}
}

func TestArea(t *testing.T) {
	result := Area(4, 2)
	expectedResult := 8.0
	if expectedResult != result {
		t.Errorf("Expected %.2f but got %.2f", expectedResult, result)
	}
}
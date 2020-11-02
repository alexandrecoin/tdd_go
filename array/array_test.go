package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sums numbers of a 5 number list", func(t *testing.T) {
		// Given
		numbersList := []int{1, 2, 3, 4, 5}
		// When
		result := Sum(numbersList)
		// Then
		expectedResult := 15
		if result != expectedResult {
			t.Errorf("Got %d but expected %d given %v", result, expectedResult, numbersList)
		}
	})

	t.Run("sums numbers of a 4 number list", func(t *testing.T) {
		// Given
		numbersList := []int{1, 2, 3, 4}
		// When
		result := Sum(numbersList)
		// Then
		expectedResult := 10
		if result != expectedResult {
			t.Errorf("Got %d but expected %d given %v", result, expectedResult, numbersList)
		}
	})
}

func TestSumAll(t *testing.T) {
	// Given
	numbersList1 := []int{1, 2, 3}
	numbersList2 := []int{1, 2}
	// When
	result := SumAll(numbersList1, numbersList2)
	// Then
	expectedResult := []int{6, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Got %v but expected %v", result, expectedResult)
	}
}

func TestSumAllTails(t *testing.T) {

	assertSumsCheck := func(t *testing.T, expectedResult, result []int) {
		t.Helper()
		if !reflect.DeepEqual(expectedResult, result) {
			t.Errorf("Expected %v but got %v", expectedResult, result)
		}
	}

	t.Run("Sum tails of non-empty slices", func(t *testing.T) {
		// Given
		list1 := []int{1, 2}
		list2 := []int{0, 9}
		// When
		result := SumAllTails(list1, list2)
		// Then
		expectedResult := []int{2, 9}
		assertSumsCheck(t, expectedResult, result)
	})

	t.Run("Sum tails of empty slices", func(t *testing.T) {
		// Given
		list1 := []int{}
		list2 := []int{0, 4, 5}
		// When
		result := SumAllTails(list1, list2)
		// Then
		expectedResult := []int{0, 9}
		assertSumsCheck(t, expectedResult, result)
	})
}

func TestSumAllItemsOfArray(t *testing.T) {
	exampleArray := []int{1, 3, 6, 7}
	got := SumAllItemsOfArray(exampleArray)
	expectedResult := 17
	if got != expectedResult {
		t.Errorf("Expected %d but got %d", expectedResult, got)
	}
}

func TestSliceArray(t *testing.T) {
	exampleSlice := []string{"Alex", "Jeannette", "Mochi", "Quiconque"}
	fromValue := 2
	toValue := 0
	result := SliceArray(exampleSlice, fromValue, toValue)

	if fromValue > toValue {
		expectedResult := []string{"Alex", "Jeannette", "Mochi", "Quiconque"}
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected %v but got %v", expectedResult, result)
		}
	}
	if fromValue < toValue {
		expectedResult := []string{"Alex", "Jeannette", "Mochi"}
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected %v but got %v", expectedResult, result)
		}
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{0, 1, 2}
	sliceToAppend := []int{3, 4, 5}
	result := CopySlice(slice, sliceToAppend)
	expected := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

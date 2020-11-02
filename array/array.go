package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

func SumAllItemsOfArray(numbersToSum []int) (result int) {
	for _, number := range numbersToSum {
		result += number
	}
	return
}

func SliceArray(array []string, fromValue int, toValue int) []string {
	if fromValue > toValue {
		return array
	}
	return array[fromValue: toValue]
}

func CopySlice(slice, sliceToAppend []int) []int {
	slice = append(slice, sliceToAppend...)
	return slice
}

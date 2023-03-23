package main

import (
	"fmt"
)

func MergeSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])

	return Merge(left, right)
}

func Merge(left []int, right []int) []int {

	result := make([]int, 0)

	i := 0
	j := 0

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	array := []int{9, 7, 5, 3, 1}
	fmt.Println("Original array:", array)

	sorted := MergeSort(array)
	fmt.Println("Sorted array:", sorted)
}

package MergeSort

import (
	"fmt"
	"math/rand"
)

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
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:len(arr)])

	return Merge(left, right)
}

func GenerateRandArr() []int {
	fmt.Print("input array length: ")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

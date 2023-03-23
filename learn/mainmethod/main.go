package main

import (
	"fmt"
	"learn/MergeSort"
)

func main() {

	arr := MergeSort.GenerateRandArr()
	fmt.Print("Original:", arr)
	sorted := MergeSort.MergeSort(arr)
	fmt.Print("Sorted array:", sorted)
}

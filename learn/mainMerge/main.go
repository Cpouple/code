package main

import (
	"fmt"
	"learn/Merge_Sort"
)

func main() {
	arr := Merge_Sort.GenerateRandArr()
	fmt.Print("Original:", arr)
	sorted := Merge_Sort.MergeSort(arr)
	fmt.Print("Sorted array:", sorted)

}

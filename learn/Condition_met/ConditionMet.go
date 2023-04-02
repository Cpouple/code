package Condition_met

import (
	"fmt"
)

func InPutNums() []int {
	fmt.Print("输入几个数")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&arr[i])
		if err != nil {
			panic(err)
		}
	}
	return arr
}
func Special() int {
	fmt.Println("输入特殊值")
	m := 0
	_, err := fmt.Scan(&m)
	if err != nil {
		panic(err)
	}
	return m
}
func JudgeSize(arr []int, m int) int {
	all := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]-arr[j] == m {
				all++
			} else if arr[j]-arr[i] == m {
				all++
			}
		}
	}

	return all
}

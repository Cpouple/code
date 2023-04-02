package main

import (
	"fmt"
	"learn/Condition_met"
)

func main() {
	arr := Condition_met.InPutNums()
	m := Condition_met.Special()
	num := Condition_met.JudgeSize(arr, m)
	fmt.Println(num)
}

package Lucky_Nums

import (
	"fmt"
)

func InPut() int {
	fmt.Println("请输入一个数：")
	nums := 0
	_, err := fmt.Scan(&nums)
	if err != nil {
		panic(err)
	}
	return nums
}
func JudgeNums(num int) int {
	fNum := num
	f := 0
	for fNum > 0 {
		f = fNum%10 + f
		fNum = fNum / 10
	}
	gNum := num
	g := 0
	for gNum > 0 {
		g = gNum&1 + g
		gNum >>= 1
	}
	all := 0
	if f == g {
		all = all + 1
	}
	return all

}
func LoopCall(num int) int {
	temp := 0
	for i := 0; i < num; i++ {
		temp += JudgeNums(i)
	}
	fmt.Println(temp)
	return temp
}

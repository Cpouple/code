package Loading_Issues

import (
	"fmt"
	"math/rand"
)

func Run() {
	weights := GenerateContainerWeight()
	c1, c2 := GenerateDockLoad(weights)
	found, res := Backtrack(weights, &c1, &c2)
	if found {
		fmt.Println("装载方案：")
		for i := 0; i < len(res); i++ {
			fmt.Printf("第 %d 艘轮船：", i+1)
			for j := 0; j < len(res[i]); j++ {
				if res[i][j] != 0 {
					fmt.Printf("%d ", j+1)
				}
			}
			fmt.Println()
		}
		fmt.Printf("第一艘轮船剩余装载量：%d\n", c1)
		fmt.Printf("第二艘轮船剩余装载量：%d\n", c2)
	} else {
		fmt.Println("无装载方案")
	}
}
func GenerateContainerWeight() []int {
	n := 0
	fmt.Print("请输入n：")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i] = rand.Intn(20) + 1 //随机生成货物质量
	}
	return weights
}
func GenerateDockLoad(weights []int) (int, int) {
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	c1Min := int(float64(totalWeight) * 0.3)
	c1Max := int(float64(totalWeight) * 0.6)
	c1 := rand.Intn(c1Max-c1Min+1) + c1Min
	c2Min := int(float64(totalWeight)) - c1
	c2Max := int(float64(totalWeight)*1.2) - c1
	c2 := rand.Intn(c2Max-c2Min+1) + c2Min
	return c1, c2 //随机生成船坞载重
}
func Backtrack(weights []int, c1, c2 *int) (found bool, res [][]int) {
	loaded := make([]bool, len(weights), len(weights))
	containerNum := 2
	res = make([][]int, containerNum, containerNum)
	for i := 0; i < containerNum; i++ {
		res[i] = make([]int, len(weights), len(weights))
	}
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(weights) {
			found = true
			return
		}
		if *c1 >= weights[i] && !loaded[i] {
			*c1 -= weights[i]
			loaded[i] = true
			res[0][i] = weights[i]
			dfs(i + 1)
			if found {
				return
			}
			*c1 += weights[i]
			loaded[i] = false
			res[0][i] = 0
		}
		if *c2 >= weights[i] && !loaded[i] {
			*c2 -= weights[i]
			loaded[i] = true
			res[1][i] = weights[i]
			dfs(i + 1)
			if found {
				return
			}
			*c2 += weights[i]
			loaded[i] = false
			res[1][i] = 0
		}
	}
	dfs(0)
	return

}

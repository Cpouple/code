package Backpack_Problem

import (
	"fmt"
	"math/rand"
)

func Test() {
	n := GetItems()
	w := GetWeights()
	Weights := make([]int, n)
	Values := make([]int, n)
	for i := 0; i < n; i++ {
		Values[i] = rand.Intn(10) + 1
		Weights[i] = rand.Intn(10) + 1
	}
	selectedItems, totalWeight, totalValue := knapsack(n, w, Weights, Values)
	// 输出结果
	fmt.Println("单个物品价值", Values)
	fmt.Println("单个物品重量", Weights)
	fmt.Printf("选中的物品编号: %v\n", selectedItems)
	fmt.Printf("背包载重量上限: %d\n", w)
	fmt.Printf("装入物品的总重量: %d\n", totalWeight)
	fmt.Printf("装入物品的总价值: %d\n", totalValue)
}
func GetItems() int {
	fmt.Print("n:")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	return n
}
func GetWeights() int {
	fmt.Print("w:")
	w := 0
	_, err := fmt.Scan(&w)
	if err != nil {
		panic(err)
	}
	return w
}
func knapsack(n, w int, weight, value []int) (selectedItems []int, totalWeight, totalValue int) {
	// 创建一个二维数组用于存储最优解
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}
	// 动态规划求解最优解
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if weight[i-1] > j {
				// 当前物品重量大于背包容量，无法放入背包
				dp[i][j] = dp[i-1][j]
			} else {
				// 考虑是否将当前物品放入背包
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i-1]]+value[i-1])
			}
		}
	}
	// 回溯查找选中的物品
	j := w
	for i := n; i > 0; i-- {
		if dp[i][j] > dp[i-1][j] {
			selectedItems = append(selectedItems, i-1)
			totalWeight += weight[i-1]
			totalValue += value[i-1]
			j -= weight[i-1]
		}
	}
	// 返回结果
	return selectedItems, totalWeight, totalValue
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

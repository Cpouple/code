package The_Longest_Common_Sub

import (
	"fmt"
	"math/rand"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Test() {
	x, y := GenerateSequences()
	// 找到这两个序列的最长公共子序列长度和具体子序列
	lcsLength, sub := longestCommonSubsequence(x, y)
	// 初始化一个二维数组存储动态规划的结果
	dp := make([][]int, len(x)+1)
	for i := range dp {
		dp[i] = make([]int, len(y)+1)
	}
	// 输出这两个序列的最长公共子序列长度和具体子序列
	fmt.Println("最长公共子序列的长度:", lcsLength)
	fmt.Println("最长的公共子序列:", sub)
}
func GenerateSequences() (x string, y string) {
	var (
		m int
		n int
	)
	fmt.Print("输入m：")
	_, err := fmt.Scan(&m)
	if err != nil {
		panic(err)
	}
	fmt.Print("输入n：")
	_, frr := fmt.Scan(&n)
	if frr != nil {
		panic(frr)
	}
	// 随机生成长度为 m 和 n 的两个字符串
	x = randString(m)
	y = randString(n)
	fmt.Printf("X: %s\nY: %s\n", x, y)
	return x, y
}

// randString 函数用于生成给定长度的随机字符串
func randString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
func longestCommonSubsequence(x, y string) (int, string) {
	m, n := len(x), len(y)
	sub := make([][][]rune, m+1) // 初始化一个空字符数组用于存储最长公共子序列
	dp := make([][]int, m+1)     // 初始化一个二维数组存储动态规划的结果
	for i := range dp {
		dp[i] = make([]int, n+1)
		sub[i] = make([][]rune, n+1)
	}
	// 使用动态规划找到这两个字符串的最长公共子序列及其长度
	for i, c1 := range x {
		for j, c2 := range y {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
				sub[i+1][j+1] = append(append([]rune{}, sub[i][j]...), c1)
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
					sub[i+1][j+1] = sub[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
					sub[i+1][j+1] = sub[i+1][j]
				}
			}
		}
	}
	// 将最长公共子序列从字符数组转换为字符串
	return dp[m][n], string(sub[m][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package Surround_area_issues

import (
	"fmt"
	"math/rand"
)

func Run() (err error) {
	var (
		n      int
		region [][]string
	)
	n, region = RegionInput()
	Printout(n, region)
	Printout(n, Solve(region))

	return
}
func Printout(n int, region [][]string) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(region[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
func RegionInput() (n int, matrix [][]string) {
	fmt.Print("请输入区域大小：")
	n = 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	var region [][]string
	region = make([][]string, n, n)
	for i := 0; i < n; i++ {
		region[i] = make([]string, n, n)
		for j := 0; j < n; j++ {
			switch rand.Intn(3) {
			case 0, 1:
				region[i][j] = "x"
			case 2:
				region[i][j] = "o"
			}
		}
	}
	return n, region
}

var n, m int

func Solve(board [][]string) [][]string {

	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		Dfs(board, i, 0)
		Dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		Dfs(board, 0, i)
		Dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == "A" {
				board[i][j] = "o"
			} else if board[i][j] == "o" {
				board[i][j] = "X"
			}
		}
	}
	return board
}

func Dfs(board [][]string, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != "o" {
		return
	}
	board[x][y] = "A"
	Dfs(board, x+1, y)
	Dfs(board, x-1, y)
	Dfs(board, x, y+1)
	Dfs(board, x, y-1)
}

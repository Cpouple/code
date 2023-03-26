package Chess_Board

import (
	"fmt"
	"strconv"
)

var (
	layer = 2
	board [][]int
)

func CoverBoard(tr, tc, dr, dc, size int) {
	// (tr,tc)表示棋盘的左上角坐标(即确定棋盘位置), (dr,dc)表示特殊方块的位置, size=2^k确定棋盘大小
	if size == 1 {
		return
	}
	layer++
	t := layer
	s := size / 2
	if dr < tr+s && dc < tc+s {
		CoverBoard(tr, tc, dr, dc, s)
	} else {
		board[tr+s-1][tc+s-1] = t
		CoverBoard(tr, tc, tr+s-1, tc+s-1, s)
	}
	if dr < tr+s && dc >= tc+s {
		CoverBoard(tr, tc+s, dr, dc, s)
	} else {
		board[tr+s-1][tc+s] = t
		CoverBoard(tr, tc+s, tr+s-1, tc+s, s)
	}
	// 判断特殊方格在不在左下棋盘
	if dr >= tr+s && dc < tc+s {
		CoverBoard(tr+s, tc, dr, dc, s)
	} else {
		board[tr+s][tc+s-1] = t
		CoverBoard(tr+s, tc, tr+s, tc+s-1, s)
	}

	// 判断特殊方格在不在右下棋盘
	if dr >= tr+s && dc >= tc+s {
		CoverBoard(tr+s, tc+s, dr, dc, s)
	} else {
		board[tr+s][tc+s] = t
		CoverBoard(tr+s, tc+s, tr+s, tc+s, s)
	}
}
func ChessBoard() int {
	fmt.Print("请输入棋盘阶数：")
	k := 0
	_, err := fmt.Scan(&k)
	if err != nil {
		panic(err)
	}
	size := 1
	for i := 0; i < k; i++ {
		size *= 2
	}
	board = make([][]int, size, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size, size)
	}
	return size
}
func SpecialPosition() (int, int) {
	fmt.Print("请输入特殊值位置：")
	row := 0
	list := 0
	_, err := fmt.Scan(&row, &list)
	if err != nil {
		panic(err)
	}
	board[row][list] = 1
	return row, list
}

func PrintBoard() {
	printCheckerboard(board)
}

func printCheckerboard(checkerboard [][]int) {
	maxInterval := 0
	currLayer := layer
	for currLayer > 0 {
		currLayer /= 10
		maxInterval++
	}
	for _, row := range checkerboard {
		for _, col := range row {
			colLength := len(strconv.Itoa(col))
			for i := 0; i < maxInterval-colLength; i++ {
				fmt.Print(" ")
			}
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

package Black_White_Chessman

import (
	"fmt"
	"strconv"
)

var (
	step = 1
)

func Chess(board []byte) {
	tail := len(board) - 2
	mid := 0
	for {
		mid = tail/2 - 1
		ChessSwap(board, mid, tail)
		tail = tail - 2
		if tail < 7 {
			break
		}
		ChessSwap(board, mid, tail)
	}
	tail++
	ChessSwap(board, mid, tail)
	mid = mid - 2
	ChessSwap(board, mid, tail)
	tail--
	ChessSwap(board, mid, tail)
	mid--
	ChessSwap(board, mid, tail)
}

func ChessSwap(board []byte, i, j int) {
	board[i], board[j] = board[j], board[i]
	board[i+1], board[j+1] = board[j+1], board[i+1]
	// print
	PrintBoard(board, step)
	step++
}

func PrintBoard(board []byte, step int) {
	fmt.Print("step " + strconv.Itoa(step) + ": ")
	for i := range board {
		fmt.Printf("%c", board[i])
	}
	fmt.Println()
}
func Chessman() (board []byte, err error) {
	fmt.Print("输入棋子个数: ")
	n := 0
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	length := 2*n + 2
	board = make([]byte, length, length)
	i := 0
	for i < n {
		board[i] = 'o'
		board[n+i] = '*'
		i++
	}
	for i < n+2 {
		board[n+i] = '-'
		i++
	}
	return
}

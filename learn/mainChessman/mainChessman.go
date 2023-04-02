package main

import "learn/Black_White_Chessman"

func main() {
	board, err := Black_White_Chessman.Chessman()
	if err != nil {
		return
	}
	Black_White_Chessman.PrintBoard(board, 0)
	Black_White_Chessman.Chess(board)
	return
}

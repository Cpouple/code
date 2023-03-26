package main

import (
	"learn/Chess_Board"
)

func main() {
	size := Chess_Board.ChessBoard()
	row, col := Chess_Board.SpecialPosition()
	Chess_Board.CoverBoard(0, 0, row, col, size)
	Chess_Board.PrintBoard()
}

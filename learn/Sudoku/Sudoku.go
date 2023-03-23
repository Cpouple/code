package main

// 判断在指定位置填入指定数字是否合法
func isValid(board [9][9]int, row int, col int, num int) bool {
	// 检查同一行是否有重复
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// 检查同一列是否有重复
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// 检查同一小方格是否有重复
	rowStart := (row / 3) * 3 // 小方格左上角所在行
	colStart := (col / 3) * 3 // 小方格左上角所在列
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

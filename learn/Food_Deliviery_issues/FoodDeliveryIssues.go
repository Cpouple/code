package Food_Deliviery_issues

import (
	"bufio"
	"fmt"
	"github.com/gogf/gf/v2/container/gqueue"
	"io"
	"os"
)

const (
	StorePoint    = 1
	CustomerPoint = 2
	ObstaclePoint = 3
)

type point struct {
	Role         int
	DeliverNeeds int
	Cost         int
}

func Run(filePath string) (err error) {
	var (
		size          int
		matrix        []point
		StoreLocation []int
		CustomerNum   int
	)
	if filePath == "" {
		matrix, StoreLocation, size, CustomerNum, err = RegionInput(os.Stdin)

	} else {
		var fileRead *os.File
		fileRead, err = os.Open(filePath)
		if err != nil {
			return
		}
		matrix, StoreLocation, size, CustomerNum, err = RegionInput(fileRead)
	}
	if err != nil {
		return
	}
	fmt.Println(bfs(matrix, size, CustomerNum, StoreLocation))

	return
}

func RegionInput(src io.Reader) (matrix []point, storeLocation []int, size, CustomerNum int, err error) {
	ObstaclesNums, StoreNums := 0, 0
	sc := bufio.NewScanner(src)
	sc.Scan()
	_, err = fmt.Sscanf(sc.Text(), "%d %d %d %d", &size, &StoreNums, &CustomerNum, &ObstaclesNums)
	if err != nil {
		return
	}
	matrix = make([]point, size*size)
	for i := 0; i < StoreNums; i++ {
		var row, col int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d", &row, &col)
		if err != nil {
			return
		}
		row--
		col--
		index := row*size + col
		matrix[index].Role = StorePoint
		storeLocation = append(storeLocation, index)
	}
	CustomerNums := CustomerNum
	for i := 0; i < CustomerNums; i++ {
		var row, col, need int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d %d", &row, &col, &need)
		if err != nil {
			return
		}
		// 坐标的开始从 (1,1) 修正为 (0,0)
		row--
		col--
		index := row*size + col
		matrix[index].Role = CustomerPoint
		// 多个客户在方格图中的同一个位置的情况
		if matrix[index].DeliverNeeds != 0 {
			CustomerNum--
		}
		matrix[index].DeliverNeeds += need
	}
	for i := 0; i < ObstaclesNums; i++ {
		var row, col int
		sc.Scan()
		_, err = fmt.Sscanf(sc.Text(), "%d %d", &row, &col)
		if err != nil {
			return
		}
		// 坐标的开始从 (1,1) 修正为 (0,0)
		row--
		col--
		matrix[row*size+col].Role = ObstaclePoint
	}
	return
}

func PrintMatrix(size int, matrix []point) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", matrix[i*size+j].Role)
		}
		fmt.Println()
	}
	fmt.Println()
}

var moveRule = [2][4]int{
	//右,下,左,上
	{0, 1, 0, -1}, // row
	{1, 0, -1, 0}} // column

func bfs(matrix []point, size, CustomerNum int, storeLocation []int) (cost int) {
	queue := gqueue.New()
	for _, index := range storeLocation {
		queue.Push(index)
	}
	for queue.Len() > 0 {
		index := queue.Pop().(int)
		for i := 0; i < 4; i++ {
			nextRow := index/size + moveRule[0][i]
			nextCol := index%size + moveRule[1][i]
			if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size {
				continue
			}
			nextIndex := nextRow*size + nextCol
			if matrix[nextIndex].Role == ObstaclePoint || matrix[nextIndex].Cost != 0 {
				continue
			}
			matrix[nextIndex].Cost = matrix[index].Cost + 1
			if matrix[nextIndex].Role == CustomerPoint {
				cost += matrix[nextIndex].Cost * matrix[nextIndex].DeliverNeeds
				CustomerNum--
			}
			if CustomerNum == 0 {
				return
			}
			queue.Push(nextIndex)
		}
	}
	return
}

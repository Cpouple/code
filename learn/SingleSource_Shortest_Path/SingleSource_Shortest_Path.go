package SingleSource_Shortest_Path

import (
	"fmt"
	"math/rand"
)

type state struct {
	Length int
	Path   []int
}

func Run() (err error) {
	size, err := VertexOfGraph()
	if err != nil {
		return
	}
	start := 0
	graph := getGraph(size, start)
	result := Dijkstra(graph, start)
	printResult(result, start)
	return
}

func VertexOfGraph() (n int, err error) {
	fmt.Print("输入图的顶点数目n：")
	_, err = fmt.Scan(&n)
	return
}

func getGraph(size int, start int) (graph [][]int) {
	graph = make([][]int, size)
	for i := 0; i < size; i++ {
		graph[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if j == start || i == j {
				graph[i][j] = -1
				continue
			}
			weight := rand.Intn(100) + 1
			if weight%3 == 0 {
				weight = -1
			}
			graph[i][j] = weight
		}
	}
	return
}
func printResult(conclusion []state, start int) {
	for i := range conclusion {

		if i == start {
			continue
		}
		fmt.Printf("开始: %d, 结束: %d, 路径长度: %d\n", start+1, i+1, conclusion[i].Length)
		PathLength := len(conclusion[i].Path)
		for j := range conclusion[i].Path {
			fmt.Print(conclusion[i].Path[j] + 1)
			if j != PathLength-1 {
				fmt.Print("->")
			}
		}
		fmt.Println()
	}
}
func Dijkstra(graph [][]int, start int) (conclusion []state) {
	conclusion = make([]state, len(graph), len(graph))
	conclusion[start] = state{
		Length: 0,
		Path:   []int{start},
	}
	// 已找出最短路径的点的集合
	haveFound := make(map[int]struct{}, len(graph))
	haveFound[start] = struct{}{}
	// 还未遍历的集合
	notTraversed := make(map[int]struct{}, len(graph))
	for i := range graph {
		if i == start {
			continue
		}
		notTraversed[i] = struct{}{}
	}
	for len(notTraversed) > 0 {
		// 更新距离
		for i := range haveFound {
			for j := range notTraversed {
				if graph[i][j] != -1 {
					if conclusion[j].Length == 0 || conclusion[i].Length+graph[i][j] < conclusion[j].Length {
						Path1 := append([]int{}, conclusion[i].Path...)
						conclusion[j] = state{
							Length: conclusion[i].Length + graph[i][j],
							Path:   append(Path1, j),
						}
					}
				}
			}
		}
		// 选出一个最短路径的点
		shortest := -1
		minIndex := -1
		for i := range notTraversed {
			if conclusion[i].Length > 0 && (shortest == -1 || conclusion[i].Length < shortest) {
				shortest = conclusion[i].Length
				minIndex = i
			}
		}
		haveFound[minIndex] = struct{}{}
		delete(notTraversed, minIndex)
	}
	return
}

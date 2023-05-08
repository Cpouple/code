package TSP_Question

import (
	"fmt"
	"math/rand"
)

const INF = 0x3f3f3f3f

type State struct {
	curPos  int      //当前访问城市
	count   int      //记录到目前为止已经访问过的城市数量。
	path    []int    //存储路径的切片
	visited [][]bool //标记每个城市是否已被访问过
}

func Run() {
	NumberOfCities()
}

func NumberOfCities() {
	var cityNum int
	fmt.Print("请输入城市数量：")
	_, err := fmt.Scan(&cityNum)
	if err != nil {
		panic(err)
	}

	// 随机生成城市之间道路长度
	distance := make([][]int, cityNum)
	for i := 0; i < cityNum; i++ {
		distance[i] = make([]int, cityNum)
		for j := 0; j < cityNum; j++ {
			if i == j {
				distance[i][j] = 0
			} else {
				distance[i][j] = rand.Intn(20) + 1
			}
		}
	}

	visited := make([]bool, cityNum)   // 记录每个城市是否已访问过
	path := make([]int, cityNum+1)     // 记录路径
	bestPath := make([]int, cityNum+1) // 记录最优路径
	minLength := -1                    // 记录最短路径

	visited[0] = true // 从第一个城市开始遍历
	TspDfs(cityNum, distance, visited, 0, 1, path, &minLength, bestPath)

	// 输出结果
	fmt.Println("各城市之间的道路长度：")
	for _, row := range distance {
		fmt.Println(row)
	}
	fmt.Printf("旅行家所走的路径（城市编号序列）：%v\n", bestPath)
	fmt.Printf("程序求出的路径长度：%d\n", minLength)

}

func TspDfs(cityNum int, distance [][]int, visited []bool, curPos int, count int, path []int, minLength *int, bestPath []int) {
	if count == cityNum { // 如果已经遍历完所有城市
		if distance[curPos][0] < INF && (distance[curPos][0]+path[count-1] < *minLength || *minLength == -1) { // 更新最短路径
			*minLength = distance[curPos][0] + path[count-1]
			copy(bestPath, path)
			bestPath[cityNum] = 0 // 最后回到起点
		}
		return
	}
	for i := 0; i < cityNum; i++ { // 遍历每个未访问过的城市
		if !visited[i] && distance[curPos][i] != INF { // 如果该城市未访问过且与当前城市有直接连接
			path[count] = distance[curPos][i]                                         // 记录走过的距离
			visited[i] = true                                                         // 标记为已访问
			TspDfs(cityNum, distance, visited, i, count+1, path, minLength, bestPath) // 继续搜索
			visited[i] = false                                                        // 回溯，将状态恢复为未访问
		}
	}
}

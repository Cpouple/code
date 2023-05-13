package TSP_Question

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"math/rand"
)

type branch struct {
	Length int    // 表示该分支所代表的路径总长度
	Path   []int  // 表示该分支所代表的经过的城市序列
	visit  []bool // 记录每个城市是否被访问过的布尔数组
}

func Test() (err error) {
	cityNum, err := NumOfCity()
	if err != nil {
		return
	}
	// 随机生成城市之间的距离矩阵
	cityDist := IntercityDistance(cityNum)
	// 输出城市距离矩阵
	CityDistance(cityDist)
	// 随机选择一个起始城市
	startCity := rand.Intn(cityNum)
	path, length := TspBfs(cityDist, startCity)
	fmt.Println("旅行家走的路径: ")
	printPath(path, length)
	return
}

// NumOfCity 随机生成城市之间的距离矩阵
func NumOfCity() (cityNum int, err error) {
	fmt.Println("请输入城市数目")
	_, err = fmt.Scan(&cityNum)
	return
}

func IntercityDistance(cityNum int) (cityMap [][]int) {
	cityMap = make([][]int, cityNum)
	for i := 0; i < cityNum; i++ {
		cityMap[i] = make([]int, cityNum)
		for j := 0; j < cityNum; j++ {
			if i == j {
				continue
			}
			cityMap[i][j] = rand.Intn(20) + 1
		}
	}
	return
}

func CityDistance(cityMap [][]int) {
	cityNum := len(cityMap)
	for i := 0; i < cityNum; i++ {
		for j := 0; j < cityNum; j++ {
			fmt.Printf("%2d ", cityMap[i][j])
		}
		fmt.Println()
	}
}

func printPath(path []int, length int) {
	pathLen := len(path)
	for i := 0; i < pathLen; i++ {
		fmt.Print(path[i] + 1)
		if i != pathLen-1 {
			fmt.Print("-")
		}
	}
	fmt.Println()
	fmt.Println(length)
}

func TspBfs(cityMap [][]int, startCity int) (path []int, length int) {
	cityNum := len(cityMap)
	// 创建一个优先队列用于存储搜索树中的分支
	queue := garray.NewSortedArray(func(a, b any) int {
		aLength := a.(*branch)
		bLength := b.(*branch)
		// 根据总长度比较两个分支
		if aLength.Length < bLength.Length {
			return -1
		}
		// 如果两个分支长度相等，则根据路径长度比较它们
		if aLength.Length == bLength.Length {
			if len(aLength.Path) > len(bLength.Path) {
				return -1
			}
			if len(aLength.Path) == len(bLength.Path) {
				return 0
			}
		}
		return 1
	})
	// 创建一个初始的分支，包含起始城市，并将其添加到优先队列中
	startB := &branch{
		Length: 0,
		Path:   []int{startCity},
		visit:  make([]bool, cityNum, cityNum),
	}
	startB.visit[startCity] = true
	queue.Add(startB)
	currLength := 0
	currPathMinLength := 0
	// 循环直到找到TSP最短路径或搜索所有可能的路径
	for queue.Len() > 0 {
		bTemp, found := queue.PopLeft()
		if !found {
			return
		}
		b := bTemp.(*branch)
		// 根据弹出的分支更新当前最长的路径和当前最短路径长度
		if b.Length > currLength {
			currLength = b.Length
			currPathMinLength = len(b.Path)
		} else if len(b.Path) < currPathMinLength {
			continue
		}
		switch len(b.Path) {
		case cityNum:
			b.Length += cityMap[b.Path[len(b.Path)-1]][startCity]
			b.Path = append(b.Path, startCity)
			queue.Add(b)
			continue
		case cityNum + 1:
			path = b.Path
			length = b.Length
			return
		}
		for i := 0; i < cityNum; i++ {
			if b.visit[i] || cityMap[b.Path[len(b.Path)-1]][i] == 0 {
				continue
			}
			nextPath := append([]int{}, b.Path...)
			nextVisitPlace := append([]bool{}, b.visit...)
			nextPath = append(nextPath, i)
			nextVisitPlace[i] = true
			queue.Add(&branch{
				Length: b.Length + cityMap[b.Path[len(b.Path)-1]][i],
				Path:   nextPath,
				visit:  nextVisitPlace,
			})
		}
	}
	return
}

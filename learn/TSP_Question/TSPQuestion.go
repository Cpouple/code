package TSP_Question

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"math/rand"
)

type embranchment struct {
	length int    // 表示该分支所代表的路径总长度
	path   []int  // 表示该分支所代表的经过的城市序列
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
	PrintPath(path, length)
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

func PrintPath(path []int, length int) {
	pathLen := len(path)
	for i := 0; i < pathLen; i++ {
		fmt.Print(path[i] + 1)
		if i != pathLen-1 {
			fmt.Print("->")
		}
	}
	fmt.Println()
	fmt.Println(length)
}

func TspBfs(cityDist [][]int, beginCity int) (path []int, length int) {
	cityNum := len(cityDist)
	// 创建一个优先队列用于存储搜索树中的分支
	array := garray.NewSortedArray(func(a, b any) int {
		aLength := a.(*embranchment)
		bLength := b.(*embranchment)
		// 根据总长度比较两个分支
		if aLength.length < bLength.length {
			return -1
		}
		// 如果两个分支长度相等，则根据路径长度比较它们
		if aLength.length == bLength.length {
			if len(aLength.path) > len(bLength.path) {
				return -1
			}
			if len(aLength.path) == len(bLength.path) {
				return 0
			}
		}
		return 1
	})
	// 创建一个初始的分支，包含起始城市，并将其添加到优先队列中
	startB := &embranchment{
		length: 0,
		path:   []int{beginCity},
		visit:  make([]bool, cityNum, cityNum),
	}
	startB.visit[beginCity] = true
	array.Add(startB)
	// 循环直到找到TSP最短路径或搜索所有可能的路径
	for array.Len() > 0 {
		bTemp, found := array.PopLeft()
		if !found {
			return
		}
		b := bTemp.(*embranchment)
		if len(b.path) == cityNum {
			b.length += cityDist[b.path[len(b.path)-1]][beginCity]
			b.path = append(b.path, beginCity)
			array.Add(b)
			continue
		}
		if len(b.path) == cityNum+1 {
			path = b.path
			length = b.length
			return
		}
		for i := 0; i < cityNum; i++ {
			if b.visit[i] || cityDist[b.path[len(b.path)-1]][i] == 0 {
				continue
			}
			nextPath := append([]int{}, b.path...)
			nextVisitPlace := append([]bool{}, b.visit...)
			nextPath = append(nextPath, i)
			nextVisitPlace[i] = true
			array.Add(&embranchment{
				length: b.length + cityDist[b.path[len(b.path)-1]][i],
				path:   nextPath,
				visit:  nextVisitPlace,
			})
		}
	}
	return
}

package Job_Scheduling_issues

import (
	"fmt"
	"math/rand"
	"sort"
)

func GetMachine() int {
	fmt.Printf("m:")
	Mac := 0
	_, err := fmt.Scan(&Mac)
	if err != nil {
		panic(err)
	}
	return Mac
}
func GetJob() int {
	fmt.Printf("n:")
	job := 0
	_, err := fmt.Scan(&job)
	if err != nil {
		panic(err)
	}
	return job
}

func scheduleJobs(jobs []int, machines int) ([][]int, int) {
	n := len(jobs)
	schedule := make([][]int, machines)  // 存储每台机器上完成的作业编号
	machineTime := make([]int, machines) // 存储每台机器完成全部作业花费的时间
	// 遍历每个作业，并将其分配给当前空闲时间最早的机器
	for i := 0; i < n; i++ {
		// 找到当前空闲时间最早的机器
		minIndex := 0
		for j := 1; j < machines; j++ {
			if machineTime[j] < machineTime[minIndex] {
				minIndex = j
			}
		}
		// 将作业分配给该机器
		schedule[minIndex] = append(schedule[minIndex], i+1)
		machineTime[minIndex] += jobs[i]
	}
	// 计算总花费时间并返回调度方案和总花费时间
	totalTime := machineTime[0]
	for i := 1; i < machines; i++ {
		if machineTime[i] > totalTime {
			totalTime = machineTime[i]
		}
	}
	return schedule, totalTime
}
func Test() {
	m := GetMachine()
	n := GetJob()
	job := make([]int, n)
	for i := 0; i < n; i++ {
		job[i] = rand.Intn(19) + 2
	}
	schedule, time := scheduleJobs(job, m)
	fmt.Printf("总花费时间：%d\n", time)
	for i := 0; i < m; i++ {
		sort.Ints(schedule[i]) // 按作业编号升序排序
		fmt.Printf("机器%d完成的作业：%v\n", i+1, schedule[i])
	}
}

package Circle_Game

import "fmt"

type kid struct {
	id   int
	next *kid
	last *kid
}

func InPutValue() (k int) {
	fmt.Print("输入特殊值k:")
	k = 0
	_, err := fmt.Scan(&k)
	if err != nil {
		panic(err)
	}
	return
}
func InPutKid() (begin *kid, n int) {
	fmt.Print("输入数字:")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	begin = &kid{id: 1}
	last := begin
	var head *kid
	for i := 2; i <= n; i++ {
		head = &kid{id: i}
		head.last = last
		last.next = head
		last = last.next
	}
	head.next = begin
	begin.last = head
	return
}
func JudgeKid(head *kid, n, k int) (winner int) {
	count := 0
	for {
		if n <= 1 {
			winner = head.id
			break
		}
		count++
		if count%10 == k || count%k == 0 {
			head.last.next = head.next
			head.next.last = head.last
			n--
		}
		head = head.next
	}
	return
}

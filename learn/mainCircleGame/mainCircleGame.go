package main

import (
	"fmt"
	"learn/Circle_Game"
)

func main() {
	begin, n := Circle_Game.InPutKid()
	k := Circle_Game.InPutValue()
	winner := Circle_Game.JudgeKid(begin, n, k)
	fmt.Println(winner)
}

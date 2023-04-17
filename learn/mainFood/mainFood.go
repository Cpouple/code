package main

import "learn/Food_Deliviery_issues"

func main() {
	err := Food_Deliviery_issues.Run("")
	if err != nil {
		return
	}
}

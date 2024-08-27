package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 60, 80, 85}
	var length = len(scores)
	var tempArray = make([]int, length+1)

	insert(scores, tempArray, 75, 2)

	scores = tempArray

	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d", scores[i])
	}
}

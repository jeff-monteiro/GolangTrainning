package main

import "fmt"

func main() {
	var scores = []int{90, 60, 40, 30, 10}
	scores = append(scores, 20)

	var length = len(scores)
	for i := 0; i < length; i++ {
		fmt.Printf("%d", scores[i])
	}

}

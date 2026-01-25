package main

import (
	"fmt"
	"slices"
)

func main() {
	var scores = []int{90, 70, 50, 60, 80, 85}
	scores = slices.Insert(scores, 2, 75)

	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d\n", scores[i])
	}
}

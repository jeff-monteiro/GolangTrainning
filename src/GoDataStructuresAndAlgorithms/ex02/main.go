package main

import "fmt"

//Anexar Tabela Linear
func append(array []int, value int) []int {
	var length = len(array)
	var tempArray = make([]int, length+1)

	for i := 0; i < length; i++ {

		tempArray[i] = array[i]
	}
	tempArray[length] = value

	return tempArray
}

func main() {
	var scores = []int{90, 70, 50, 60, 80, 85}
	scores = append(scores, 75)

	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d", scores[i])
	}
}

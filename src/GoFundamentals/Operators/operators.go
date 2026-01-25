package main

import "fmt"

func main() {
	// Arithmetics Operators
	soma := 10 + 2
	subtracao := 10 - 2
	multi := 10 * 2
	divisao := 10 / 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multi, divisao, restoDaDivisao)
	atribuitionOperators()
	relationalOperators()
	logicalOperators()
	unaryOperators()
}

// Atribuition Operators
func atribuitionOperators() {
	var variavel string = "My name"
	variavel2 := "This is Us"

	fmt.Println(variavel, variavel2)
}

// Relational Operators
func relationalOperators() {
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
}

// Logical Operators
func logicalOperators() {
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}

// Unary Operators
func unaryOperators() {
	temp := 10
	temp++     // temp = temp + 1
	temp += 15 // temp = temp + 15
	temp--     // temp = temp - 1

	temp *= 3 // temp = temp * 3
	temp /= 2 // temp = temp / 2
	temp %= 2 // temp = temp % 2
	fmt.Println(temp)
}

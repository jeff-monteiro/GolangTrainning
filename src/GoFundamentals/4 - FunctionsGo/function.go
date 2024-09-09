package main

import "fmt"

func main() {
	name := "John Doe"
	age := 25
	sayHello(name, age)
	printOut(12)
	dataPrint()

	// Anonymous function
	var f = func() {
		fmt.Println("Hello World!")
	}
	f()

	resultadoSoma, resultadoSubt := calcMath(10, 15)
	fmt.Println(resultadoSoma, resultadoSubt)
}

func sayHello(name string, age int) {
	fmt.Println(name, age)
}

func printOut(numberLogs int) {
	fmt.Println(numberLogs)
}

func dataPrint() {
	data := 10
	for i := 0; i <= data; i++ {
		fmt.Println(i)
	}
}

// Return multiple values
func calcMath(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

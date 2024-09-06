package main

import "fmt"

func main() {
	name := "John Doe"
	age := 25
	sayHello(name, age)
	printOut(12)
	dataPrint()

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

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!")
		time.Sleep(100 * time.Millisecond)
	}
}

func sayTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("Just Testing Goroutines")
		time.Sleep(150 * time.Millisecond)
	}
}

func Goroutines() {
	go sayHello()
	go sayTest()
	fmt.Println("Você se garantiu!")
	time.Sleep(1 * time.Second) // Inicia a goroutine

}

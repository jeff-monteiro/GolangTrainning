package main

import (
	"fmt"
	"time"
)

func Channels() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("Finalizou a escrita no canal!")
	}()

	time.Sleep(time.Second * 1)
	for valor := range ch {
		fmt.Println("Leitura do canal:", valor)
	}

}

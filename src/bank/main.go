package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroConta   int
	numeroAgencia int
	saldoConta    float64
}

func main() {
	contaJefferson := ContaCorrente{"Jefferson", 589, 256934, 3000.05}
	contaJoelma := ContaCorrente{"Joelma", 123, 568784, 7000.01}
	fmt.Println(contaJefferson)
	fmt.Println(contaJoelma)

}

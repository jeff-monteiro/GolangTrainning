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

	//Other way to use Struct
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroConta = 876687
	contaDaCris.numeroAgencia = 001
	contaDaCris.saldoConta = 1000.23

	fmt.Println(contaDaCris)
}

package main

import (
	"fmt"

	"bank/contas"
)

func main() {
	contaJefferson := contas.ContaCorrente{Titular: "Jefferson", SaldoConta: 300}
	contaMaria := contas.ContaCorrente{Titular: "Maria", SaldoConta: 200}

	status := contaJefferson.Transferir(100, &contaMaria)
	fmt.Println(status)
	fmt.Println(contaJefferson)
	fmt.Println(contaMaria)

}

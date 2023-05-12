package main

import (
	"fmt"

	"bank/contas"
)

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	contaDenis.Sacar(4000)

	fmt.Println(contaDenis.ObterSaldo())
}

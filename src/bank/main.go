package main

import (
	"fmt"

	"bank/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	PagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.ObterSaldo())

	contaLuiza := contas.ContaCorrente{}
	contaLuiza.Depositar(600)
	PagarBoleto(&contaLuiza, 400)

	fmt.Println(contaLuiza.ObterSaldo())
}

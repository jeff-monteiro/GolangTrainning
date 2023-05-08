package main

import (
	"fmt"

	"bank/clientes"
	"bank/contas"
)

func main() {
	contaJeff := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Jefferson",
		CPF:       "123.222.111.33",
		Profissão: "Desenvolvedor"},
		NumeroAgencia: 325698, NumeroConta: 123, SaldoConta: 200}

	fmt.Println(contaJeff)

}

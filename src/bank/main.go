package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroConta   int
	numeroAgencia int
	saldoConta    float64
}

// Calculate withdraw from TPM
func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldoConta
	if podeSacar {
		c.saldoConta -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saque não realizado"
	}
}

// Calculate deposit from TPM || função com multiplos retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldoConta += valorDoDeposito
		return "Deposito feito com sucesso!", c.saldoConta
	} else {
		return "O valor do depósito menor que zero!", c.saldoConta
	}

}

// Make transfer between accounts
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldoConta && valorDaTransferencia > 0 {
		c.saldoConta -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaJefferson := ContaCorrente{titular: "Jefferson", saldoConta: 300}
	contaMaria := ContaCorrente{titular: "Maria", saldoConta: 200}

	status := contaJefferson.Transferir(100, &contaMaria)
	fmt.Println(status)
	fmt.Println(contaJefferson)
	fmt.Println(contaMaria)

}

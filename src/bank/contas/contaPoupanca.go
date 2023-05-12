package contas

import (
	"bank/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldoConta                           float64
}

// Calculate withdraw from TPM
func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldoConta
	if podeSacar {
		c.saldoConta -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saque não realizado"
	}
}

// Calculate deposit from TPM || função com multiplos retornos
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldoConta += valorDoDeposito
		return "Deposito feito com sucesso!", c.saldoConta
	} else {
		return "O valor do depósito menor que zero!", c.saldoConta
	}

}

// Allow get balance
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldoConta
}

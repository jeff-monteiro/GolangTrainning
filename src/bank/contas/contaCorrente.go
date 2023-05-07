package contas

import "bank/clientes"

// Using composition on struct(it's used struct inside struct)
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroConta   int
	NumeroAgencia int
	SaldoConta    float64
}

// Calculate withdraw from TPM
func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.SaldoConta
	if podeSacar {
		c.SaldoConta -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saque não realizado"
	}
}

// Calculate deposit from TPM || função com multiplos retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.SaldoConta += valorDoDeposito
		return "Deposito feito com sucesso!", c.SaldoConta
	} else {
		return "O valor do depósito menor que zero!", c.SaldoConta
	}

}

// Make transfer between accounts
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.SaldoConta && valorDaTransferencia > 0 {
		c.SaldoConta -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

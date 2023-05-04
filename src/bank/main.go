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

func main() {
	contaJefferson := ContaCorrente{}
	contaJefferson.titular = "Jefferson"
	contaJefferson.saldoConta = 800

	fmt.Println(contaJefferson.saldoConta)

	fmt.Println(contaJefferson.Sacar(-100))
	fmt.Println(contaJefferson.saldoConta)

}

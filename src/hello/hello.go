package main

import "fmt"

func main() {
	name := "Jefferson Monteiro"
	version := 1.1
	fmt.Println("Faaaala galera! Meu nome é ", name)
	fmt.Println("A versão do GO é a ", version)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando) //That's a reference variable
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Iniciando Monitoramento...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Não conheço esse comando")
	}
}

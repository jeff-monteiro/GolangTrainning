package main

import "fmt"

func main() {

	showIntro()
	showMenu()

	// if comando == 1 {
	// 	fmt.Println("Iniciando Monitoramento...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }
	comando := readCommand()
	switch comando {
	case 1:
		fmt.Println("Iniciando Monitoramento...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não conheço esse comando")
	}
}

// Modularizando minha aplicação
func showIntro() {
	name := "Jefferson Monteiro"
	version := 1.1
	fmt.Println("Faaaala galera! Meu nome é ", name)
	fmt.Println("A versão do GO é a ", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command) //That's a reference variable
	fmt.Println("O comando escolhido foi", command)
	return command
}

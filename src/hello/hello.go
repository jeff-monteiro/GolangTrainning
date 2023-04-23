package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorizing = 5
const delay = 5

func main() {

	showIntro()
	for {
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
			startMonitorizing()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
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

func startMonitorizing() {
	fmt.Println("Iniciando Monitoramento...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://caelum.com.br"}
	fmt.Println(sites)
	for i := 0; i < monitorizing; i++ {
		for i, site := range sites {
			fmt.Println(i, site)
			testSiteUp(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSiteUp(site string) {
	respo, _ := http.Get(site)

	if respo.StatusCode == 200 {
		fmt.Println("Site", site, "carregou com sucesso...OK!")
	} else {
		fmt.Println("Error!", site, "site não pôde ser carregado! StatusCode:", respo.StatusCode)
	}
}

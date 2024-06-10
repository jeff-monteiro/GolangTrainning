package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorizing = 3
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
			printOutLog()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

// Modularizing my application
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
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://caelum.com.br"}
	sites := readFileSite()
	for i := 0; i < monitorizing; i++ {
		for i, site := range sites {
			fmt.Println(i, ":", site)
			testSiteUp(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSiteUp(site string) {
	respo, err := http.Get(site)

	if err != nil {
		fmt.Println("Page Not Found", err)
	}
	if respo.StatusCode == 200 {
		fmt.Println("Site", site, "carregou com sucesso...OK!")
		registerLogs(site, true)
	} else {
		fmt.Println("Error!", site, "site não pôde ser carregado! StatusCode:", respo.StatusCode)
		registerLogs(site, false)
	}
}

func readFileSite() []string {
	var sites []string

	file, err := os.Open("File.txt")
	// file, err := ioutil.ReadFile("File.txt")

	if err != nil {
		fmt.Println("Error!", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registerLogs(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online " + strconv.FormatBool(status) + "\n")
	file.Close()

}

func printOutLog() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}

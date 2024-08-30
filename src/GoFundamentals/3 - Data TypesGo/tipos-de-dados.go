package main

import (
	"errors"
	"fmt"
)

// Data Types
func main() {

	// INT INICIO
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias == apelido
	//RUNE é um alias para int32
	//BYTE é um alias para uint8

	//INT FIM

	//FLOAT INICIO

	var number3 float32 = 123.45
	fmt.Println(number3)

	var number4 float64 = 12365.89
	fmt.Println(number4)

	//FLOAT FIM

	//STRING INICIO

	var nomeVariavel1 string = "Nome Da Variavel1"
	fmt.Println(nomeVariavel1)

	var nomeVariavel2 = "Nome Da Variavel2"
	fmt.Println(nomeVariavel2)

	char := 'B'
	fmt.Println(char)

	//STRING FIM

	//BOOLEAN INICIO

	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)

	//DATA TYPE ERROR

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}

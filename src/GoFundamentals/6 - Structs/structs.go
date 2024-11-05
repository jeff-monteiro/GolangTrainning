package main

import "fmt"

type person struct {
	nome  string
	idade int
	endereco
}

type endereco struct {
	endereco string
}

func main() {
	fmt.Println("That's an example of struct")

	// Declaring a struct using variable declaration
	var usuario1 person
	usuario1.nome = "John Doe"
	usuario1.idade = 30
	usuario1.endereco = endereco{"Rua: Maria Padilha, 35"}
	fmt.Println(usuario1)

	// Declaring a struct using type inference
	usuario2 := person{"Jeff Doe", 33, endereco{"Rua dos bobos, 0"}}
	fmt.Println(usuario2)
}

package main

/*func main() {
	gavetas := []string{}

	gavetas = append(gavetas, "golang", "java", "python", "javascript")
	fmt.Println(gavetas[0:3])
}*/

/*
func main() {
	var armario = map[string]int64{}
	fmt.Println(armario)
	armario["golang"] = 1
	armario["java"] = 2
	armario["puthon"] = 3
	armario["javascript"] = 4

	fmt.Println(armario)

	delete(armario, "puthon")
	fmt.Println(armario)

	if name, ok := armario["puthon"]; ok {
		fmt.Println("Existe na lista", name, ok)

	} else {
		fmt.Println("Não existe na lista")
	}

}
*/

/*
func main() {

	fmt.Println("Qual o dia da semana?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("É hoje!")
	case today + 1:
		fmt.Println("É amanhã!")
	case today + 2:
		fmt.Println("Será em dois dias!")
	default:
		fmt.Println("Tá é longe ainda!")
	}

}
*/

/*
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p1 := Pessoa{"Katita", 29}
	fmt.Println(p1)

	var p2 *Pessoa = &p1
	fmt.Println(p2)

	p2 = &Pessoa{"Jeff", 30}
	p1.Nome = "Karina Morais"
	p1.Idade = 35
	fmt.Println(p1)
	fmt.Println(p2)

}
*/

/*

"*************************** DON'T FORGET THIS SUBJECT **************************"

"SEE SOMETHING ABOUT GOROUTINES"
"NEXT TOPIC IS CHANNELS IN GOLANG"

*/

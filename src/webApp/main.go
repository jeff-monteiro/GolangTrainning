package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Product{
		{"Camisa", "Vermelha Ferrari", 29.90, 8},
		{"Blusão", "Preto com capus", 129.90, 7},
		{"Tenis", "Adidas Premium", 230.30, 4},
		{"Relógio", "Mormaii Preto Fosco", 500.80, 2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)

}

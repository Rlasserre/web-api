package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta", Description: "Azul, estampada", Price: 39, Quantity: 15},
		{"Tenis", "Comfortavel", 89, 3},
		{"Fone", "Muito bom", 18, 8},
		{"Blank", "...", 0, 999},
	}

	templates.ExecuteTemplate(w, "Index", products)
}

package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func dbconnection() *sql.DB {
	connectionStr := "user=postgres dbname=gostore password=241917 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	db := dbconnection()
	defer db.Close()
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

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
	Id, Quantity      int
	Price             float64
}

var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	db := dbconnection()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbconnection()

	selectAllProducts, err := db.Query("Select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	templates.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}

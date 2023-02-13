package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-api/models"
)

var templates = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.ScanProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloatConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro, conversão do preço.")
		}

		quantityIntConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro, conversão da quantidade.")
		}

		models.CreateNewProduct(name, description, priceFloatConverted, quantityIntConverted)
	}
	http.Redirect(w, r, "/", 301)
}

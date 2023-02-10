package main

import (
	"html/template"
	"net/http"
	"web-api/models"
)

var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.ScanProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)

}

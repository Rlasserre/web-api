package controllers

import (
	"html/template"
	"net/http"
	"web-api/models"
)

var templates = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.ScanProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)

}

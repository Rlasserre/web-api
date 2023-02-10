package main

import (
	"net/http"
	"web-api/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}

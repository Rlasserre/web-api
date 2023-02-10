package routes

import (
	"net/http"
	"web-api/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
}

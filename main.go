package main

import (
	"net/http"

	"github.com/edusant/cursoGo/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.GetUsers)
	http.ListenAndServe(":8080", nil)
}

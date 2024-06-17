package main

import (
	"net/http"

	"github.com/yourusername/project/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.HelloHandler)
	http.ListenAndServe(":8080", nil)
}

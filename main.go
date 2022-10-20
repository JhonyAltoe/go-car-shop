package main

import (
	"fmt"
	"net/http"

	routes "github.com/jhony/go-car-shop/api/routes"
)

func main() {
	routes := routes.CarRoutes()
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", routes)
}

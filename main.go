package main

import (
	"net/http"

	mongoconfig "github.com/jhony/go-car-shop/api/database/config"
	car_router "github.com/jhony/go-car-shop/api/routes"
)

func main() {
	mongoconfig.Connect()
	routes := car_router.CarRoutes()
	http.ListenAndServe(":8080", routes)
}

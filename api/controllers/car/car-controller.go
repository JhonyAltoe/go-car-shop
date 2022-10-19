package car_controller

import (
	"net/http"

	carService "github.com/jhony/go-car-shop/api/services/car"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarController interface {
	GetAll(res *http.ResponseWriter, req *http.Request)
}

type carController struct {
	CarService carService.TCarService
}

func New(collection *mongo.Collection) TCarController {
	return &carController{
		CarService: carService.New(collection),
	}
}

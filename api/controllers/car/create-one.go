package car_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhony/go-car-shop/api/database/entities"
)

func (c *carController) CreateOne(res http.ResponseWriter, req *http.Request) {
	var car *entities.TCar
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		res.Write([]byte(err.Error()))
	}

	createdCar, _ := c.CarService.CreateOne(context.Background(), car)

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	defer json.NewEncoder(res).Encode(*createdCar)
}

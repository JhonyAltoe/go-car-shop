package car_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhony/go-car-shop/api/database/entities"
)

func (c carController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	var car *entities.TCar
	err := json.NewDecoder(req.Body).Decode(&car)
	param := mux.Vars(req)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	updatedCar, err := c.CarService.Update(context.Background(), param["id"], car)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(updatedCar)
}

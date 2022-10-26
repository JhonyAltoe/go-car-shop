package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (c carController) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	var car *entities.TCar
	err := json.NewDecoder(req.Body).Decode(&car)
	param := mux.Vars(req)
	if err != nil {
		customErr := helpers.CustomErrBuilder(
			http.StatusBadRequest,
			"Wrong json body",
			"Controller.Update",
			err,
		)
		helpers.SendError(customErr, res)
		return
	}

	updatedCar, customErr := c.CarService.Update(context.Background(), param["id"], car)
	if customErr != nil {
		helpers.SendError(customErr, res)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(updatedCar)
}

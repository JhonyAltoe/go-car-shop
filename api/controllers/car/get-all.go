package car_controller

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *carController) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	cars, err := c.CarService.GetAll(context.Background())
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(cars)
}

package car_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *carController) GetOne(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	car, err := c.CarService.GetOne(context.Background(), param["id"])
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(car)
}

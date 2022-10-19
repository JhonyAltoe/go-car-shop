package car_controller

import (
	"context"
	"encoding/json"
	"net/http"
)

// func YourHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Gorilla!\n"))
// }

func (c *carController) GetAllCars(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	cars, err := c.CarService.GetAll(context.Background())
	if err != nil {
		utils.SendError(err, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(cars)
}

package car_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c carController) DeleteOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	d, err := c.CarService.DeleteOne(context.Background(), param["id"])
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`"{ message: "` + err.Error() + `" }`))
	}

	var message = map[string]string{"message": ""}

	if d.DeletedCount == 0 {
		message["message"] = "this car doen't exist"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(message)
		return
	}

	message["message"] = "deleted"
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(message)
}

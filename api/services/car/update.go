package car_service

import (
	"context"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (s carService) Update(ctx context.Context, id string, car *entities.TCar) (*entities.TCar, *helpers.CustomError) {
	updatedCar, err := s.carModel.Update(ctx, id, car)
	if err != nil {
		return updatedCar, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			"Internal server error",
			"Services.Update",
			err,
		)
	}

	return updatedCar, nil
}

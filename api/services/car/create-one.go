package car_service

import (
	"context"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
	v "github.com/jhonyaltoe/go-car-shop/api/validations"
)

func (c *carService) CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, *helpers.CustomError) {
	customErr := v.Validate(car)

	if customErr != nil {
		return &entities.TCar{}, customErr
	}
	carResponse, err := c.carModel.CreateOne(ctx, car)
	if err != nil {
		return &entities.TCar{}, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			err.Error(),
			"Services.CreateOne.CarResponse",
			err,
		)
	}
	return carResponse, nil
}

package car_service

import (
	"context"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (m *carService) GetOne(ctx context.Context, id string) (*entities.TCar, *helpers.CustomError) {
	car, err := m.carModel.GetOne(&ctx, id)
	if err != nil {
		return car, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			"this car doesn't exist",
			"Services.GetOne",
			err,
		)
	}
	return car, nil
}

package car_service

import (
	"context"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (m *carService) GetAll(ctx context.Context) (*[]entities.TCar, *helpers.CustomError) {
	c, err := m.carModel.GetAll(ctx)
	if err != nil {
		return &c, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			"Internal server error",
			"Services.GetAll",
			err,
		)
	}
	return &c, nil
}

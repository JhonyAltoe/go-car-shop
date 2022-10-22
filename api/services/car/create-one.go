package car_service

import (
	"context"

	"github.com/jhony/go-car-shop/api/database/entities"
)

func (c *carService) CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, error) {
	return c.carModel.CreateOne(ctx, car)
}

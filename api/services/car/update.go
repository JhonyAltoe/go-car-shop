package car_service

import (
	"context"

	"github.com/jhony/go-car-shop/api/database/entities"
)

func (s carService) Update(ctx context.Context, id string, car *entities.TCar) (*entities.TCar, error) {
	updatedCar, err := s.carModel.Update(ctx, id, car)
	if err != nil {
		return &entities.TCar{}, err
	}

	return updatedCar, nil
}

package car_service

import (
	"context"

	"github.com/jhony/go-car-shop/api/database/entities"
)

func (m *carService) GetAll(ctx context.Context) (*[]entities.TCar, error) {
	c, err := m.carModel.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

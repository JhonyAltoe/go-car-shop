package car_service

import (
	"context"
	"fmt"

	"github.com/jhony/go-car-shop/api/database/entities"
)

func (m *carService) GetOne(ctx context.Context, id string) (*entities.TCar, error) {
	car, err := m.carModel.GetOne(&ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return car, nil
}

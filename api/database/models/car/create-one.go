package car_model

import (
	"context"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carModel) CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, error) {
	car.ID = primitive.NewObjectID()
	_, err := c.collection.InsertOne(ctx, &car)
	if err != nil {
		return &entities.TCar{}, err
	}
	return car, nil
}

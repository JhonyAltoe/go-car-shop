package car_model

import (
	"context"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carModel) GetOne(ctx *context.Context, id string) (*entities.TCar, error) {
	var car *entities.TCar
	mongoId, _ := primitive.ObjectIDFromHex(id)
	err := c.collection.FindOne(*ctx, bson.M{"_id": mongoId}).Decode(&car)
	if err != nil {
		return &entities.TCar{}, err
	}
	return car, nil
}

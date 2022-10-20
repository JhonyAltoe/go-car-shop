package car_model

import (
	"context"
	"fmt"

	"github.com/jhony/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carModel) GetAll(ctx context.Context) ([]entities.TCar, error) {
	var cars []entities.TCar
	cur, err := c.collection.Find(ctx, primitive.M{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = cur.All(ctx, &cars)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return cars, nil
}

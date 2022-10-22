package car_model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *carModel) DeleteOne(ctx *context.Context, id string) (*mongo.DeleteResult, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return &mongo.DeleteResult{}, err
	}
	v, err := c.collection.DeleteOne(*ctx, bson.M{"_id": mongoId})
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	return v, nil
}

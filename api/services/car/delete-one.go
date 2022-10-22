package car_service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s carService) DeleteOne(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	v, err := s.carModel.DeleteOne(&ctx, id)
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	return v, nil
}

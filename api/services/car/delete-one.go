package car_service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s carService) DeleteOne(ctx context.Context, id string) (*mongo.DeleteResult, *helpers.CustomError) {
	r, err := s.carModel.DeleteOne(&ctx, id)
	if err != nil {
		return r, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			"Internal server error",
			"Services.DeleteOne",
			err,
		)
	}
	if r.DeletedCount == 0 {
		return r, helpers.CustomErrBuilder(
			http.StatusInternalServerError,
			"this car doesn't exist",
			"Services.DeleteOne",
			errors.New("the datadase could not find the 'id' car"),
		)
	}
	return r, nil
}

package car_model

import (
	"context"

	"github.com/jhonyaltoe/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarModel interface {
	GetAll(ctx context.Context) ([]entities.TCar, error)
	CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, error)
	GetOne(ctx *context.Context, id string) (*entities.TCar, error)
	DeleteOne(ctx *context.Context, id string) (*mongo.DeleteResult, error)
	Update(ctx context.Context, id string, car *entities.TCar) (*entities.TCar, error)
}

type carModel struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) TCarModel {
	return &carModel{
		collection,
	}
}

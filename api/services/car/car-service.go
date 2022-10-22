package car_service

import (
	"context"

	"github.com/jhony/go-car-shop/api/database/entities"
	carModel "github.com/jhony/go-car-shop/api/database/models/car"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarService interface {
	GetAll(ctx context.Context) ([]entities.TCar, error)
	CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, error)
}

type carService struct {
	carModel carModel.TCarModel
}

func New(collection *mongo.Collection) TCarService {
	return &carService{
		carModel.New(collection),
	}
}

package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type TCar struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Model    string             `json:"model,omitempty" bson:"model,omitempty"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty"`
	Color    string             `json:"color,omitempty" bson:"color,omitempty"`
	Status   bool               `json:"status,omitempty" bson:"status,omitempty"`
	BuyValue int                `json:"buyValue,omitempty" bson:"buyValue,omitempty"`
	DoorsQty int                `json:"doorsQty,omitempty" bson:"doorsQty,omitempty"`
	SeatsQty int                `json:"seatsQty,omitempty" bson:"seatsQty,omitempty"`
}

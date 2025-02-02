package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID         primitive.ObjectID `json:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

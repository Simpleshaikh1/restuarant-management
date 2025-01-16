package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Menu_id    string             `json:"menu_id"`
}

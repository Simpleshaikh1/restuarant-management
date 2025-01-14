package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name" validate:"required"`
	Category   string             `bson:"category" validate:"required"`
	Start_Date time.Time          `bson:"start_date"`
	End_Date   time.Time          `bson:"end_date"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Menu_id    string             `bson:"menu_id"`
}

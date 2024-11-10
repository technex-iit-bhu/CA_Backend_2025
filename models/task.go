package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" binding:"required"`
	Points      int                `json:"points,omitempty" bson:"points,omitempty" binding:"required"`
	Deadline    time.Time          `json:"deadline,omitempty" bson:"deadline,omitempty" binding:"required"`
	ImageUrl    string             `json:"image_url,omitempty" bson:"image_url,omitempty" binding:"required"`
}
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TaskSubmission struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task         string             `json:"task,omitempty" bson:"task,omitempty" binding:"required"`
	User         string             `json:"user,omitempty" bson:"user,omitempty" binding:"required"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
	Timestamp    time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty" binding:"required"`
	DriveLink    string             `json:"drive_link,omitempty" bson:"drive_link,omitempty" binding:"required"`
	ImageUrl     string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Verified     bool               `json:"verified,omitempty" bson:"verified,omitempty"`
	AdminComment string             `json:"admin_comment,omitempty" bson:"admin_comment,omitempty"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CA_ID          string             `json:"ca_id,omitempty" bson:"ca_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Username       string             `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
	PhoneNumber    string             `json:"phone,omitempty" bson:"phone,omitempty" binding:"required"`
	WhatsappNumber string             `json:"whatsapp,omitempty" bson:"whatsapp,omitempty" binding:"required"`
	Institute      string             `json:"institute,omitempty" bson:"institute,omitempty" binding:"required"`
	City           string             `json:"city,omitempty" bson:"city,omitempty"`
	Postal_address string             `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
	Pin_code       string             `json:"pin_code,omitempty" bson:"pin_code,omitempty"`
	WhyChooseYou   string             `json:"why_choose_you,omitempty" bson:"why_choose_you,omitempty"`
	IsChosen       bool               `json:"is_chosen,omitempty" bson:"is_chosen,omitempty"`
	WereCA         bool               `json:"were_ca,omitempty" bson:"were_ca,omitempty"`
	Points         int                `json:"points,omitempty" bson:"points,omitempty"`
	Year           int                `json:"year,omitempty" bson:"year,omitempty"`
	Branch         string             `json:"branch,omitempty" bson:"branch,omitempty"`
	Tasks          map[string]bool    `json:"completed_tasks,omitempty" bson:"completed_tasks,omitempty"`
	ReferralCode   string             `json:"referral_code,omitempty" bson:"referral_code,omitempty" binding:"required"`
	ReferralCount  int                `json:"referral_count,omitempty" bson:"referral_count,omitempty" binding:"required"`
	IsReferred     bool               `json:"is_referred,omitempty" bson:"is_referred,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" binding:"required"`
	UpdatedAt      time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" binding:"required"`
}

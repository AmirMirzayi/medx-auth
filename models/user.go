package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserName    string             `bson:"username" json:"username" binding:"required"`
	FirstName   string             `bson:"first_name" json:"first_name" binding:"required"`
	Password    string             `bson:"password" json:"password" binding:"required"`
	Address     string             `bson:"address" json:"address,omitempty"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number" binding:"required"`
}

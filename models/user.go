package models

type User struct {
	Name        string `json:"name" bson:"user_name"`
	Email       string `json:"email" bson:"user_email"`
	PhoneNumber int    `json:"phone_number" bson:"user_phone_number"`
}

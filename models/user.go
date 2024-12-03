package models

type User struct {
	Id          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"user_name"`
	Email       string `json:"email" bson:"user_email"`
	PhoneNumber int    `json:"phone_number" bson:"user_phone_number"`
}

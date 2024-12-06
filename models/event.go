package models

type Event struct {
	Title       string `json:"title" bson:"event_title"`
	Description string `json:"description" bson:"event_description"`
	Date        string `json:"date" bson:"event_date"`
	Location    string `json:"location" bson:"event_location"`
}

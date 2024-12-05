package services

import "github.com/danielanugr/GatherGo-EventTracker/models"

type EventService interface {
	CreateEvent(*models.Event) error
	GetEventById(*string) (*models.Event, error)
	GetAll() (*[]models.Event, error)
	UpdateEvent(*models.Event) error
	DeleteEvent(*string) error
}

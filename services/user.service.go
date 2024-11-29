package services

import "github.com/danielanugr/GatherGo-EventTracker/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUserById(*string) (*models.User, error)
	GetAll() (*[]models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

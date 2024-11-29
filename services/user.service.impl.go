package services

import (
	"context"
	"github.com/danielanugr/GatherGo-EventTracker/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}

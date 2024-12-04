package services

import (
	"context"
	"errors"
	"github.com/danielanugr/GatherGo-EventTracker/models"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUserById(id *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(u.ctx)

	err = cursor.All(u.ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	query := bson.D{bson.E{Key: "id", Value: user.Id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: user.Id}, bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_email", Value: user.Email}, bson.E{Key: "user_phone_number", Value: user.PhoneNumber}}}}
	result, _ := u.userCollection.UpdateOne(u.ctx, query, update)
	if result.MatchedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	query := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := u.userCollection.DeleteOne(u.ctx, query)
	if result.DeletedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

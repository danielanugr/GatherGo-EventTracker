package main

import (
	"context"
	"fmt"
	"github.com/danielanugr/GatherGo-EventTracker/controllers"
	"github.com/danielanugr/GatherGo-EventTracker/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"time"
)

var (
	server         *gin.Engine
	userService    services.UserService
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
	cancel         context.CancelFunc
)

func init() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	userCollection = mongoClient.Database("gathergo").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/api")
	userController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":8080"))
}

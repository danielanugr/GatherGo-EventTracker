package controllers

import (
	"github.com/danielanugr/GatherGo-EventTracker/models"
	"github.com/danielanugr/GatherGo-EventTracker/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(UserService services.UserService) UserController {
	return UserController{
		UserService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.UserService.GetUserById(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/:id", uc.GetUserById)
	userRoute.GET("/", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/:id/delete", uc.CreateUser)
}

package controllers

import (
	"github.com/danielanugr/GatherGo-EventTracker/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func newUserController(userService services.UserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/:id", uc.GetUserById)
	userRoute.GET("/", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete", uc.CreateUser)
}

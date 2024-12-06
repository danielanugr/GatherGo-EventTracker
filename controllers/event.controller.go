package controllers

import (
	"net/http"

	"github.com/danielanugr/GatherGo-EventTracker/models"
	"github.com/danielanugr/GatherGo-EventTracker/services"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService services.EventService
}

func NewEventController(eventService services.EventService) *EventController {
	return &EventController{
		eventService,
	}
}

func (ec *EventController) CreateEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.eventService.CreateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func (ec *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := ec.eventService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (ec *EventController) GetEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := ec.eventService.GetEventById(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (ec *EventController) UpdateEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.eventService.UpdateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (ec *EventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ec.eventService.DeleteEvent(&id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}

func (ec *EventController) RegisterEventRoutes(rg *gin.RouterGroup) {
	eventRoute := rg.Group("/event")
	eventRoute.POST("/", ec.CreateEvent)
	eventRoute.GET("/", ec.GetAllEvents)
	eventRoute.GET("/:id", ec.GetEventById)
	eventRoute.PUT("/:id", ec.UpdateEvent)
	eventRoute.DELETE("/:id", ec.DeleteEvent)
}

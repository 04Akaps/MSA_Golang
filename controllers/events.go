package controllers

import (
	"fmt"
	"net/http"

	"GO_MSA/middleware"
	"GO_MSA/models"
	"GO_MSA/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService services.Event
}

type ReqTest struct {
	NameTwo   string `json:"name_two" binding:"required"`
	NameThree string `json:"name_three" binding:"required"`
}

func NewEventController(event services.Event) EventController {
	return EventController{
		EventService: event,
	}
}

type FindByIdReq struct {
	Id string `uri:"id" binding:"required"`
}

func (ec *EventController) FindEventById(ctx *gin.Context) {
	var req FindByIdReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ec.EventService.FindEventByName(req.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "FindEventById Failed", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type FindByNameReq struct {
	Name string `uri:"name" binding:"required"`
}

func (ec *EventController) FindEventByName(ctx *gin.Context) {
	var req FindByNameReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ec.EventService.FindEventByName(req.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "FindEventByName Failed", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventController) AllEvents(ctx *gin.Context) {
	// FindAllAvaliableEvents() (*[]models.EventModel, error)

	result, err := ec.EventService.FindAllAvaliableEvents()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Find ALl Failed", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (ec *EventController) NewEvent(ctx *gin.Context) {
	var req models.EventModel

	bodyCheckError := middleware.CheckBodyBinding(&req, ctx)

	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError, "status": -1})
		return
	}

	result, err := ec.EventService.AddEvent(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "DB 저장 실패..", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, (result))
}

func (ec *EventController) RegisterEventRoutes(server *gin.Engine) {
	eventpath := server.Group("/events")
	eventpath.Use(gin.CustomRecovery(func(ctx *gin.Context, rec interface{}) {
		fmt.Println("panic이 일어 날 떄만 동작 하는 middleWare")
		fmt.Println(rec) // rec에서는 panic에서 넘어오는 값이 적히게 된다.
		// 후에 가능하다면 middleWare에 따로 정리할 예정
	}))

	eventpath.GET("/findById/:id", ec.FindEventById)
	eventpath.GET("/findByName/:name", ec.FindEventByName)
	eventpath.GET("/", ec.AllEvents)
	eventpath.POST("/", ec.NewEvent)
}

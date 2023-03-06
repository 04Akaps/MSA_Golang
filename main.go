package main

import (
	"context"
	"fmt"
	"log"

	"GO_MSA/config"
	"GO_MSA/controllers"
	"GO_MSA/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var server *gin.Engine

var (
	// event variable
	eventCtx        context.Context
	eventService    services.Event
	eventController controllers.EventController
)

var envConfig config.Config

func init() {
	envConfig = config.LoadConfig(".")
	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	gin.EnableJsonDecoderDisallowUnknownFields()

	trustAddress := []string{"http://127.0.0.1"} // nginx를 사용할 예정이기 떄문에

	server = gin.New()

	server.SetTrustedProxies(trustAddress)

	server.Use(gin.Logger())   // Logger를 통해서 DefaultWriter에 로그를 기록 -> 로그 형태 변환
	server.Use(gin.Recovery()) // panic이 발생하면 500에러

	// Set Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = trustAddress // 모든 도메인에 대한 요청을 허용
	server.Use(cors.New(config))
	// Set Event Controller
	eventCtx = context.Background()
	eventService = services.NewEventService(eventCtx)
	eventController = controllers.NewEventController(eventService)

	eventpath := server.Group("/events")
	eventpath.Use(gin.CustomRecovery(func(ctx *gin.Context, rec interface{}) {
		fmt.Println("panic이 일어 날 떄만 동작 하는 middleWare")
		fmt.Println(rec) // rec에서는 panic에서 넘어오는 값이 적히게 된다.
	}))
	eventController.RegisterEventRoutes(eventpath)
}

func main() {
	log.Fatal(server.Run(envConfig.ServerAddress))
}

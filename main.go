package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"GO_MSA/config"
	"GO_MSA/controllers"
	"GO_MSA/mongo"
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

var ctxMongo context.Context

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
	ctxMongo = context.Background()
}

func main() {
	// mongo Session

	mongoDBLayout, err := mongo.NewMongoSession(ctxMongo, envConfig)
	if err != nil {
		log.Fatal("Mongo Session Connection Error", err)
	}

	err = mongoDBLayout.Session.Ping(ctxMongo, nil)

	if err != nil {
		log.Fatal("mongo Connection ERror ping", err)
	}

	eventService = services.NewEventService(eventCtx, mongoDBLayout)
	eventController = controllers.NewEventController(eventService)

	eventpath := server.Group("/events")
	eventpath.Use(gin.CustomRecovery(func(ctx *gin.Context, rec interface{}) {
		fmt.Println("panic이 일어 날 떄만 동작 하는 middleWare")
		fmt.Println(rec) // rec에서는 panic에서 넘어오는 값이 적히게 된다.
		// 후에 가능하다면 middleWare에 따로 정리할 예정
	}))
	eventController.RegisterEventRoutes(eventpath)

	tlsConfig, err := config.GetTlsConfig(envConfig)
	if err != nil {
		log.Fatal("GetTlsConfig Error", err)
	}

	// HTTPS 설정
	httpsServer := &http.Server{
		Addr:      envConfig.ServerAddress,
		Handler:   server,
		TLSConfig: tlsConfig,
	}

	err = httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("Server Start Error", err)
	}

	defer mongoDBLayout.Session.Disconnect(ctxMongo) // 리소스를 줄이기 위해서 mongo에 대한 Close를 defer로 호출
}

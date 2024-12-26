package main

import (
	"ProjectX/api/cmd/route"
	"fmt"

	_ "ProjectX/api/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Controller
// @version         1.0
// @description     API Server for Bot Application
// @BasePath  /
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	e := echo.New()

	// Добавьте CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // Домен вашего фронтенда
		AllowMethods: []string{echo.GET, echo.POST},     // Разрешённые методы
		AllowHeaders: []string{echo.HeaderContentType},  // Разрешённые заголовки
	}))

	e.GET("/swagger*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{fmt.Sprintf("http://%s:%s/swagger/doc.json", "127.0.0.1", "8080")}
	}))

	botAPI := e.Group("/botapi/v1")
	route.RoutesBot(botAPI)
	front := e.Group("/front/v1")
	route.RoutesFront(front)

	e.Logger.Fatal(e.Start(":8080"))
}

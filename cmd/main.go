package main

import (
	"ProjectX/api/cmd/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Добавьте CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // Домен вашего фронтенда
		AllowMethods: []string{echo.GET, echo.POST},     // Разрешённые методы
		AllowHeaders: []string{echo.HeaderContentType},  // Разрешённые заголовки
	}))

	botAPI := e.Group("/botapi/v1")
	route.RoutesBot(botAPI)
	front := e.Group("/front/v1")
	route.RoutesFront(front)

	e.Logger.Fatal(e.Start(":8080"))
}

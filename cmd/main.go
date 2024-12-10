package main

import (
	botapihandlers "ProjectX/api/cmd/botAPIhandlers"
	fronthandlers "ProjectX/api/cmd/frontHandlers"

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

	// Группа маршрутов для botAPI
	botAPI := e.Group("/botapi/v1")
	botAPI.POST("/commands", botapihandlers.HandleCommand)

	// Группа маршрутов для фронтенда
	front := e.Group("/front/v1")
	front.POST("/auth", fronthandlers.HandleAuth)

	// Запуск сервера на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}

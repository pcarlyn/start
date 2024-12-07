package main

import (
	botapihandlers "ProjectX/api/cmd/botAPIhandlers"
	fronthandlers "ProjectX/api/cmd/frontHandlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	botAPI := e.Group("/botapi/v1")
	botAPI.POST("/commands", botapihandlers.HandleCommand)
	front := e.Group("/front/v1")
	front.POST("/auth", fronthandlers.HandleAuth)

	e.Logger.Fatal(e.Start(":8080"))
}

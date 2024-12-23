package route

import (
	botapihandlers "ProjectX/api/cmd/botAPIhandlers"

	"github.com/labstack/echo/v4"
)

func RoutesBot(group *echo.Group) {
	group.POST("/commands", botapihandlers.HandleCommand)
}

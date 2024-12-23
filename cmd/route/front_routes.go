package route

import (
	fronthandlers "ProjectX/api/cmd/frontHandlers"

	"github.com/labstack/echo/v4"
)

func RoutesFront(group *echo.Group) {
	group.POST("/auth", fronthandlers.HandleAuth)
}

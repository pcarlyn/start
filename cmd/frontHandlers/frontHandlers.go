package fronthandlers

import (
	"ProjectX/api/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleAuth(c echo.Context) error {

	var auth models.Auth
	var result = "false"

	if err := c.Bind(&auth); err != nil {
		return err
	}

	if auth.Login == "admin" && auth.Password == "admin" {
		result = "true"
	}

	return c.JSON(http.StatusOK, result)
}

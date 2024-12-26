package fronthandlers

import (
	"ProjectX/api/internal/models"
	"ProjectX/api/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleAuth(c echo.Context) error {

	var auth models.Auth

	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusInternalServerError, models.UserProfile{})
	}
	resp, status := utils.GetAuth(models.RepositoryURL, "auth", auth.Login)

	switch status {
	case http.StatusNotFound:
		return c.JSON(http.StatusNotFound, models.UserProfile{})
	case http.StatusOK:
		if auth.Password != resp.PswdHache {
			return c.JSON(http.StatusForbidden, auth)
		}
	}

	return c.JSON(status, auth)
}

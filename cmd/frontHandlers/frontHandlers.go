package fronthandlers

import (
	"ProjectX/api/internal/models"
	"ProjectX/api/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler for all comands front
// @Summary Summary
// @Description Description
// @Tags front
// @Accept json
// @Produce json
// @Param data body models.Auth true "Authorization"
// @Success 200 {object} models.Auth "Success Response"
// @Failure 403 {object} models.ErrorResponse "Server error"
// @Failure 404 {object} models.ErrorResponse "Server error"
// @Failure 500 {object} models.ErrorResponse "Server error"
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

package botapihandlers

import (
	"ProjectX/api/internal/models"
	"ProjectX/api/internal/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleCommand(c echo.Context) error {
	var userData models.Response
	var update = false
	var answer string

	if err := c.Bind(&userData); err != nil {
		return err
	}
	_, err := utils.GetUserById("http://localhost:8088", "users", userData.Result[0].Message.From.ID)
	if err != nil {
		fmt.Printf("error sending GETuserById request: %v\n", err)
		_, err := utils.PostUser("http://localhost:8088", "users", userData)
		if err != nil {
			fmt.Printf("error sending POSTuser request: %v\n", err)
		}
		update = true
		answer = "first"
	}
	if !update {
		_, err = utils.PatchUserById("http://localhost:8088", "users", userData)
		if err != nil {
			fmt.Printf("error sending PATCHuserById request: %v", err)
		}
		answer = "second"
	}
	result, err := utils.GetAnswer("http://localhost:8088", "bot/commands/", "start", answer)
	if err != nil {
		fmt.Printf("error sending GETanswer request: %v\n", err)
	}
	return c.JSON(http.StatusOK, result)
}

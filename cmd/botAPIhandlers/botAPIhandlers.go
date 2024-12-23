package botapihandlers

import (
	"ProjectX/api/internal/models"
	"ProjectX/api/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleCommand(c echo.Context) error {
	var userData models.Result
	var answer models.Response
	if err := c.Bind(&userData); err != nil {
		return err
	}
	_, statusCode := utils.GetUserById(models.RepositoryURL, "users", userData.Message.From.ID)
	switch statusCode {
	case http.StatusOK:
		exp, _ := utils.GetAnswer(models.RepositoryURL, "commands", "start", "registred")
		answer = models.Response{
			ID:     uint(userData.Message.Chat.ID),
			ChatID: userData.Message.Chat.ID,
			Text:   exp,
		}
	case http.StatusNotFound:
		_, statusCode := utils.GetProfileById(models.RepositoryURL, "profile", userData.Message.From.ID)
		switch statusCode {
		case http.StatusOK:
			exp, _ := utils.GetAnswer(models.RepositoryURL, "commands", "start", "registred")
			answer = models.Response{
				ID:     uint(userData.Message.Chat.ID),
				ChatID: userData.Message.Chat.ID,
				Text:   exp,
			}
		case http.StatusNotFound:
			exp, _ := utils.GetAnswer(models.RepositoryURL, "commands", "start", "registred")
			answer = models.Response{
				ID:     uint(userData.Message.Chat.ID),
				ChatID: userData.Message.Chat.ID,
				Text:   exp,
			}
		}
	}
	// get porile (id user)
	//  if err {
	//    profile
	// 	post profile (id user)
	//  }
	//  init user
	//  tg - api
	//  user
	//  post user
	// fmt.Printf("error sending GETuserById request: %v\n", err)

	return c.JSON(http.StatusOK, answer)
}

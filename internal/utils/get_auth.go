package utils

import (
	"ProjectX/api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAuth(baseURL, command, login string) (models.UserProfile, int) {

	url := fmt.Sprintf("%s/%s/%s", baseURL, command, login)

	resp, err := http.Get(url)
	if err != nil {
		return models.UserProfile{}, resp.StatusCode
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return models.UserProfile{}, resp.StatusCode
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.UserProfile{}, http.StatusInternalServerError
	}
	var result models.UserProfile

	if err := json.Unmarshal(body, &result); err != nil {
		return models.UserProfile{}, http.StatusInternalServerError
	}

	return result, http.StatusOK
}

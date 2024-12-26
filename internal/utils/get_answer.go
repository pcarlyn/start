package utils

import (
	"ProjectX/api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetAnswer(baseURL, route, command, commandType string) (models.Answer, int) {
	var result models.Answer
	fullurl := fmt.Sprintf("%s/%s", baseURL, route)

	queryParams := url.Values{}
	queryParams.Add("command", command)
	queryParams.Add("type", commandType)

	resp, err := http.Get(fullurl + "?" + queryParams.Encode())
	if err != nil {
		return models.Answer{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.Answer{}, http.StatusNotFound
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Answer{}, http.StatusInternalServerError
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.Answer{}, http.StatusInternalServerError
	}
	return result, http.StatusOK
}

package utils

import (
	"ProjectX/api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetProfileById(baseURL, route string, id int) (models.Result, int) {
	url := fmt.Sprintf("%s/%s/%d", baseURL, route, id)

	resp, err := http.Get(url)
	if err != nil {
		return models.Result{}, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return models.Result{}, resp.StatusCode
	}
	if resp.StatusCode == http.StatusNotFound {
		return models.Result{}, http.StatusNotFound
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Result{}, http.StatusInternalServerError
	}

	var result models.Result
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Result{}, http.StatusInternalServerError
	}

	return result, http.StatusOK
}

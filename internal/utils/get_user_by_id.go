package utils

import (
	"ProjectX/api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserById(baseURL, route string, id int) (models.Result, int) {

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

// func PatchUserById(baseURL, route string, updateData models.Result) (models.Result, error) {

// 	url := fmt.Sprintf("%s/%s/%d", baseURL, route, updateData.Message.From.ID)

// 	jsonData, err := json.Marshal(updateData)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error marshalling request body: %v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error creating request: %v", err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error reading models.Result body: %v", err)
// 	}

// 	if err := json.Unmarshal(respBody, &updateData); err != nil {
// 		return models.Result{}, fmt.Errorf("error unmarshalling models.Result body: %v", err)
// 	}
// 	return updateData, nil
// }

// func PostUser(baseURL, route string, userData models.Result) (models.Result, error) {

// 	url := fmt.Sprintf("%s/%s", baseURL, route)

// 	jsonData, err := json.Marshal(userData)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error marshalling request body: %v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error creating request: %v", err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return models.Result{}, fmt.Errorf("error reading models.Result body: %v", err)
// 	}

// 	if err := json.Unmarshal(respBody, &userData); err != nil {
// 		return models.Result{}, fmt.Errorf("error unmarshalling models.Result body: %v", err)
// 	}
// 	return userData, nil
// }

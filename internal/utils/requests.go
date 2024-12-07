package utils

import (
	"ProjectX/api/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetAnswer(baseURL, route, command, commandType string) (string, error) {
	var result string
	fullurl := fmt.Sprintf("%s/%s", baseURL, route)

	queryParams := url.Values{}
	queryParams.Add("command", command)
	queryParams.Add("type", commandType)

	resp, err := http.Get(fullurl + "?" + queryParams.Encode())
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading models.Response body: %v", err)
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("error unmarshalling models.Response body: %v", err)
	}
	return result, nil
}

func PatchUserById(baseURL, route string, updateData models.Response) (models.Response, error) {

	url := fmt.Sprintf("%s/%s/%d", baseURL, route, updateData.Result[0].Message.From.ID)

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		return models.Response{}, fmt.Errorf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.Response{}, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Response{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, fmt.Errorf("error reading models.Response body: %v", err)
	}

	if err := json.Unmarshal(respBody, &updateData); err != nil {
		return models.Response{}, fmt.Errorf("error unmarshalling models.Response body: %v", err)
	}
	return updateData, nil
}

func PostUser(baseURL, route string, userData models.Response) (models.Response, error) {

	url := fmt.Sprintf("%s/%s", baseURL, route)

	jsonData, err := json.Marshal(userData)
	if err != nil {
		return models.Response{}, fmt.Errorf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.Response{}, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Response{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, fmt.Errorf("error reading models.Response body: %v", err)
	}

	if err := json.Unmarshal(respBody, &userData); err != nil {
		return models.Response{}, fmt.Errorf("error unmarshalling models.Response body: %v", err)
	}
	return userData, nil
}

func GetUserById(baseURL, route string, id int) (models.Response, error) {

	url := fmt.Sprintf("%s/%s/%d", baseURL, route, id)

	resp, err := http.Get(url)
	if err != nil {
		return models.Response{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return models.Response{}, fmt.Errorf("error sending request: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Response{}, fmt.Errorf("error reading models.Response body: %v", err)
	}

	var result models.Response
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Response{}, fmt.Errorf("error unmarshalling models.Response body: %v", err)
	}
	return result, nil

}

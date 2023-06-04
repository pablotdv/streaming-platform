package restreamer

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(login LoginRequest) (LoginResponse, error) {
	jsonData, err := json.Marshal(login)
	if err != nil {
		return LoginResponse{}, err
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return LoginResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LoginResponse{}, err
	}
	defer resp.Body.Close()

	var loginResponse LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		return LoginResponse{}, err
	}

	return loginResponse, nil
}

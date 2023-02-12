package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ApiKey string
}

const baseURL = "https://api.hypixel.net"

func NewClient(ApiKey string) *User {
	return &User{ApiKey: ApiKey}
}

type getPlayer struct {
	Username string `json:"username"`
}

func (s *User) GetPlayer(username string) (*getPlayer, error) {
	url := fmt.Sprintf(baseURL+"/%s/todos/%d", s.ApiKey, username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data getPlayer
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *User) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.ApiKey, _)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

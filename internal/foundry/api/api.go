package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type FoundryAPI struct {
	endpoint    string
	acess_token string
}

/*
route should be formatted with an intiial /
*/
func CreateFoundryAPI(route string) *FoundryAPI {

	accessToken := os.Getenv("FOUNDRY_TOKEN")
	if accessToken == "" {
		log.Fatal("FOUNDRY_ACCESS_TOKEN is required")
		return nil
	}
	endpoint := os.Getenv("FOUNDRY_ENDPOINT")
	if accessToken == "" {
		log.Fatal("FOUNDRY_ACCESS_TOKEN is required")
		return nil
	}

	return &FoundryAPI{
		acess_token: accessToken + route,
		endpoint:    endpoint,
	}

}

func (api *FoundryAPI) apply(payload map[string]interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, api.endpoint, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+api.acess_token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil

}

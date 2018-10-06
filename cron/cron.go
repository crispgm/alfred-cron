package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	// APIURL ...
	APIURL = "https://cronexpressiondescriptor.azurewebsites.net/api/descriptor/?locale=en-US&expression=%s"
)

var (
	// ErrAPIFailure ...
	ErrAPIFailure = errors.New("Call API failed")
	// ErrJSONFailure ...
	ErrJSONFailure = errors.New("Decode JSON failed")
)

// APIResponse for crontab API
type APIResponse struct {
	Description string `json:"description"`
}

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

// Call api
func Call(cronExpr string) (string, error) {
	url := fmt.Sprintf(APIURL, url.QueryEscape(cronExpr))
	resp, err := client.Get(url)
	if err != nil {
		return "", ErrAPIFailure
	}
	defer resp.Body.Close()
	var apiResp APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return "", ErrJSONFailure
	}
	return apiResp.Description, err
}

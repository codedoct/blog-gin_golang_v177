package test

import (
	"encoding/json"
	"fmt"

	"blog-gin_golang_v177/lib/httpcli"
	"blog-gin_golang_v177/lib/response"
)

var (
	baseURL = "https://staging-loyaltyservice.mysuperindo.co.id"
)

func HttpCallGet(url string, req map[string]string) (*response.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	httpClient := httpcli.HttpClientParam{
		Params:  req,
		Headers: headers,
		URL:     fmt.Sprintf("%v/%v", baseURL, url),
	}
	resp, err := httpcli.Get(&httpClient)
	if err != nil {
		return nil, err
	}
	stringResp := fmt.Sprintf("%v", resp)

	var respData response.Response
	err = json.Unmarshal([]byte(stringResp), &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

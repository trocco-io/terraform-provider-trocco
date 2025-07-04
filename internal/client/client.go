package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	version2 "terraform-provider-trocco/version"
)

const (
	RegionJapan  = "japan"
	RegionIndia  = "india"
	RegionKorea  = "korea"
	BaseURLJapan = "https://trocco.io"
	BaseURLIndia = "https://in.trocco.io"
	BaseURLKorea = "https://kr.trocco.io"
)

var RegionBaseURLMap = map[string]string{
	RegionJapan: BaseURLJapan,
	RegionIndia: BaseURLIndia,
	RegionKorea: BaseURLKorea,
}

type TroccoClient struct {
	BaseURL    string
	APIKey     string
	httpClient *http.Client
}

func NewTroccoClient(apiKey string) *TroccoClient {
	return &TroccoClient{
		BaseURL:    BaseURLJapan,
		APIKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func NewTroccoClientWithRegion(apiKey, region string) (*TroccoClient, error) {
	baseURL, ok := RegionBaseURLMap[region]
	if !ok {
		return nil, fmt.Errorf("invalid region: %s", region)
	}
	return &TroccoClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

func NewDevTroccoClient(apiKey, baseURL string) *TroccoClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &TroccoClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		httpClient: httpClient,
	}
}

type ErrorOutput struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (client *TroccoClient) do(
	method, path string, input interface{}, output interface{},
) error {
	var reqBody io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(b)
	}

	url := client.BaseURL + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Token "+client.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "terraform-provider-trocco "+version2.ProviderVersion)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode >= 400 {
		var errorOutput ErrorOutput
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(respBody, &errorOutput)
		if err != nil {
			return fmt.Errorf("%s", resp.Status)
		}

		if errorOutput.Error != "" {
			return fmt.Errorf("%s", errorOutput.Error)
		}
		if errorOutput.Message != "" {
			return fmt.Errorf("%s", errorOutput.Message)
		}
	}
	if output == nil {
		return nil
	}
	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, output)
	if err != nil {
		return err
	}
	return nil
}

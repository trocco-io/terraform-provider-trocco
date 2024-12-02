package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Error string `json:"error"`
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
	// return fmt.Errorf("%+v", reqBody)

	url := client.BaseURL + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Token "+client.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
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
		return fmt.Errorf("%s", errorOutput.Error)
	}
	if output == nil {
		return nil
	}
	respBody, err := io.ReadAll(resp.Body)

	// return fmt.Errorf("%+v", string(respBody))
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, output)
	if err != nil {
		return err
	}
	return nil
}

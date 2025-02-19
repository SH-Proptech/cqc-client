package cqc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	pathProviders = "/providers"
	pathLocations = "/locations"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiToken   string
}

type Config struct {
	APIToken string
	Timeout  time.Duration
}

// NewClient initializes a new API client with a timeout.
func NewClient(cfg *Config) *Client {
	if cfg == nil {
		cfg = &Config{
			APIToken: getEnvOrPanic("CQC_API_PASSWORD"),
			Timeout:  60 * time.Second,
		}
	}

	return &Client{
		httpClient: &http.Client{Timeout: cfg.Timeout},
		baseURL:    "https://api.service.cqc.org.uk/public/v1",
		apiToken:   cfg.APIToken,
	}
}

// NewRequest creates an HTTP request with authentication headers.
func (c *Client) NewRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	fullURL := c.baseURL + endpoint

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", c.apiToken)

	return req, nil
}

// Get sends a GET request and unmarshals the response into a generic type.
func Get[T any](c *Client, path string) (*T, error) {
	return request[T](c, "GET", path, nil)
}

// request performs an HTTP request and decodes the response.
func request[T any](c *Client, method, path string, body io.Reader) (*T, error) {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read response body: %v", err)
	// }
	// fmt.Printf("Raw JSON Response: %s\n", string(bodyBytes))
	// resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &result, nil
}

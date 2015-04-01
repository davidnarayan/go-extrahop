// Package extrahop is a Go client for the ExtraHop REST API
package extrahop

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client is an ExtraHop API client
type Client struct {
	Host   string
	ApiKey string
	Scheme string

	httpClient *http.Client
}

const (
	DefaultScheme = "https"
	DefaultPath   = "/api/1.0"
)

// Create a new ExtraHop API client
func NewClient(host, apikey string) *Client {
	// Setup a custom transport that ignores SSL certificate errors.
	// TODO: This probably isn't what needs to happen long term...
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Client{
		Host:       host,
		ApiKey:     apikey,
		Scheme:     DefaultScheme,
		httpClient: &http.Client{Transport: tr},
	}
}

// Post sends a POST request and returns the decoded response
func (c *Client) Post(path string, body io.Reader) (f interface{}, err error) {
	return c.do("POST", path, body)
}

// Get sends a GET request and returns the decoded response
func (c *Client) Get(path string) (f interface{}, err error) {
	return c.do("GET", path, nil)
}

// do sends an HTTP request and returns the response
func (c *Client) do(method, path string, data io.Reader) (f interface{}, err error) {
	url := fmt.Sprintf("%s://%s%s%s", c.Scheme, c.Host, DefaultPath, path)
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, fmt.Errorf("unable to create request for %s: %s", url, err)
	}

	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Add the API key for authentication
	// TODO: Allow this to be set dynamically from the /api-docs path
	req.Header.Set("X-api-key", c.ApiKey)

	// Send the request
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("request error from %s: %s", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response error from %s: %s", url,
			resp.Status)
	}

	// Parse and decode body
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("unable to decode response from %s: %s", url, err)
	}

	return f, nil
}

// Explore returns the set of available API paths
/*
func (c *Client) Explore() {
	f, err := c.Get("/api-docs")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
*/

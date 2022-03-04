package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
)

type ClientError string

func (e ClientError) Error() string {
	return string(e)
}

func NewClient(httpClient HTTPClient, baseURL string) (*Client, error) {
	return &Client{
		http:    httpClient,
		baseURL: baseURL,
	}, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	http    HTTPClient
	baseURL string
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	return req, err
}

func GetEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		log.Println("value", value)
		return value
	}
	log.Println("env key", os.Getenv(key))
	log.Println("default", def)
	return def
}
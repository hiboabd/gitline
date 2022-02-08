package handlers

import (
	"fmt"
	"github.com/joho/godotenv"
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

	personalAccessToken := GetEnvVariable("PERSONAL_ACCESS_TOKEN")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", personalAccessToken))

	return req, err
}

func GetEnvVariable(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		return os.Getenv(key)
	}
}
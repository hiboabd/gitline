package handlers

import (
	"io"
	"net/http"
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

	req.Header.Add("OPG-Bypass-Membrane", "1")

	return req, err
}


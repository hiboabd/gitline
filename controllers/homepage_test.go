package controllers

import (
	"github.com/hiboabd/gitline/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRenderHomepage(t *testing.T) {
	mockClient := &mocks.MockClient{}
	client, _ := NewClient(mockClient, "http://mock-server:2000")

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       nil,
		}, nil
	}

	template, expectedResponse, err := client.RenderHomepage()
	assert.Equal(t, template, "home.gotmpl")
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, expectedResponse)
}

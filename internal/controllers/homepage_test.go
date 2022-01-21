package controllers

import (
	"github.com/hiboabd/gitline/internal/handlers"
	"github.com/hiboabd/gitline/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderHomepage(t *testing.T) {
	assertions := assert.New(t)

	mockClient := &mocks.MockClient{}
	client, _ := handlers.NewClient(mockClient, "http://mock-server:2000")
	template := &mockTemplates{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/path", nil)

	handler := renderHomepage(client, template)
	err := handler(w, r)

	assertions.Nil(err)

	resp := w.Result()
	assertions.Equal(http.StatusOK, resp.StatusCode)
}
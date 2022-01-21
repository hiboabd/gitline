package controllers

import (
	"github.com/hiboabd/gitline/internal/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUsernameInformation struct {
	count      int
	err        error
	UserRepositories handlers.RepositoryData
}

func (m *mockUsernameInformation) GetRepositoryData(username string) (handlers.RepositoryData, error) {
	m.count += 1
	return m.UserRepositories, m.err
}

func TestNavigateToTimeline(t *testing.T) {
	assertions := assert.New(t)

	client := &mockUsernameInformation{}
	template := &mockTemplates{}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/path", nil)

	handler := renderTimeline(client, template)
	err := handler(w, r)

	assertions.Nil(err)

	resp := w.Result()
	assertions.Equal(http.StatusOK, resp.StatusCode)
}


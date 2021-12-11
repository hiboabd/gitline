package controllers

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetRepositoryData(t *testing.T) {
	r := getRouter(true)
	r.HTMLRender = CreateMyRender("../web/templates")

	r.GET("/timeline", GetRepositoryData)

	req, _ := http.NewRequest("GET", "/timeline", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<h1>Timeline</h1>") > 0

		return statusOK && pageOK
	})
}

func SetUpTestData() Repositories {
	repositories := Repositories{
		Repository{
			ID:   1,
			Name: "Repository 1",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl.com",
			CreatedAt: "12/05/2020",
			UpdatedAt: "13/05/2020",
			PushedAt:  "13/05/2020",
			Size:      70,
			Language:  "JavaScript",
		},
		Repository{
			ID:   2,
			Name: "Repository 2",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl2.com",
			CreatedAt: "12/04/2020",
			UpdatedAt: "13/04/2020",
			PushedAt:  "13/04/2020",
			Size:      100,
			Language:  "Python",
		},
		Repository{
			ID:   3,
			Name: "Repository 3",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl3.com",
			CreatedAt: "22/10/2021",
			UpdatedAt: "05/12/2021",
			PushedAt:  "05/12/2021",
			Size:      30,
			Language:  "Golang",
		},
	}
	return repositories
}

func TestFormatRepositories(t *testing.T) {
	unformattedRepository1 := apiRepository{}
	unformattedRepository1.ID = 1
	unformattedRepository1.Name = "Test Repository"
	unformattedRepository1.Owner = Owner{ID: 1}
	unformattedRepository1.URL = "Test URL"
	unformattedRepository1.CreatedAt = "2020-05-12T17:25:38Z"
	unformattedRepository1.UpdatedAt = "2020-05-12T17:25:38Z"
	unformattedRepository1.PushedAt = "2020-05-12T17:25:38Z"
	unformattedRepository1.Size = 40
	unformattedRepository1.Language = "Golang"

	unformattedRepositoryList := RepositoriesList{unformattedRepository1}
	unformattedData := apiRepositories{
		unformattedRepositoryList,
	}

	expectedResponse := Repositories{
		Repository{
			ID:   1,
			Name: "Test Repository",
			Owner: Owner{
				1,
			},
			URL:       "Test URL",
			CreatedAt: "12/05/2020",
			UpdatedAt: "12/05/2020",
			PushedAt:  "12/05/2020",
			Size:      40,
			Language:  "Golang",
		},
	}

	assert.Equal(t, expectedResponse, formatRepositories(unformattedData))
}

func TestSortRepositoriesByCreatedDate(t *testing.T) {
	testData := SetUpTestData()

	expectedData := Repositories{
		Repository{
			ID:   2,
			Name: "Repository 2",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl2.com",
			CreatedAt: "12/04/2020",
			UpdatedAt: "13/04/2020",
			PushedAt:  "13/04/2020",
			Size:      100,
			Language:  "Python",
		},
		Repository{
			ID:   1,
			Name: "Repository 1",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl.com",
			CreatedAt: "12/05/2020",
			UpdatedAt: "13/05/2020",
			PushedAt:  "13/05/2020",
			Size:      70,
			Language:  "JavaScript",
		},
		Repository{
			ID:   3,
			Name: "Repository 3",
			Owner: Owner{
				1,
			},
			URL:       "www.testurl3.com",
			CreatedAt: "22/10/2021",
			UpdatedAt: "05/12/2021",
			PushedAt:  "05/12/2021",
			Size:      30,
			Language:  "Golang",
		},
	}
	assert.Equal(t, sortRepositoriesByCreatedDate(testData), expectedData)
}

func TestFormatStringToDateObject(t *testing.T) {
	testData := "2021-12-08T17:25:38Z"
	expectedResponse, _ := time.Parse(time.RFC3339, testData)
	assert.Equal(t, expectedResponse, formatStringToDateObject(testData))
}

func TestFormatDate(t *testing.T) {
	testData := "2021-12-08T17:25:38Z"
	expectedResponse := "08/12/2021"
	assert.Equal(t, expectedResponse, formatDate(testData))
}

package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//func TestGetRepositoryData(t *testing.T) {
//	r := getRouter(true)
//
//	r.GET("/timeline", GetRepositoryData)
//
//	req, _ := http.NewRequest("GET", "/timeline", nil)
//
//	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
//		statusOK := w.Code == http.StatusOK
//
//		p, err := ioutil.ReadAll(w.Body)
//		pageOK := err == nil && strings.Index(string(p), "<h1>Timeline</h1>") > 0
//
//		return statusOK && pageOK
//	})
//}

func TestFormatRepositories(t *testing.T) {

}

func TestSortRepositoriesByCreatedDate(t *testing.T) {


}

func TestFormatStringToDateObject(t *testing.T) {
	testData := "2021-12-08T17:25:38Z"
	expectedResponse , _ := time.Parse(time.RFC3339, testData)
	assert.Equal(t, expectedResponse, formatStringToDateObject(testData))
}

func TestFormatDate(t *testing.T) {
	testData := "2021-12-08T17:25:38Z"
	expectedResponse := "08/12/2021"
	assert.Equal(t, expectedResponse, formatDate(testData))
}
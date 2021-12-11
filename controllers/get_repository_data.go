package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Owner struct {
	ID int `json:"id"`
}

type Repository struct {
	ID        int
	Name      string
	Owner     Owner
	URL       string
	CreatedAt string
	UpdatedAt string
	PushedAt  string
	Size      int
	Language  string
}

type Repositories []Repository

type apiRepository struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Owner     Owner  `json:"owner"`
	URL       string `json:"html_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt  string `json:"pushed_at"`
	Size      int    `json:"size"`
	Language  string `json:"language"`
}

type RepositoriesList []apiRepository

type apiRepositories struct {
	RepositoriesList `json:"repositories"`
}

const userFacingDateFormat string = "02/01/2006"

func GetRepositoryData(c *gin.Context) {
	apiUrl := getEnv("API_URL", "")
	resp, err := http.Get(apiUrl + "/api/v1/repositories")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result apiRepositories
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	userRepositories := formatRepositories(result)

	c.HTML(
		http.StatusOK,
		"timeline",
		gin.H{
			"Repositories": userRepositories,
		},
	)
}

func formatRepositories(r apiRepositories) Repositories {
	var repositories Repositories
	for _, t := range r.RepositoriesList {
		var repository = Repository{
			ID:   t.ID,
			Name: t.Name,
			Owner: Owner{
				t.Owner.ID,
			},
			URL:       t.URL,
			CreatedAt: formatDate(t.CreatedAt),
			UpdatedAt: formatDate(t.UpdatedAt),
			PushedAt:  formatDate(t.PushedAt),
			Size:      t.Size,
			Language:  t.Language,
		}

		repositories = append(repositories, repository)
	}

	sortedRepositories := sortRepositoriesByCreatedDate(repositories)
	return sortedRepositories
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}

func sortRepositoriesByCreatedDate(repositories Repositories) Repositories {
	// sorted in ascending date order
	sort.Slice(repositories, func(i, j int) bool {
		dateOne, _ := time.Parse(userFacingDateFormat, repositories[i].CreatedAt)
		dateTwo, _ := time.Parse(userFacingDateFormat, repositories[j].CreatedAt)
		return dateOne.Before(dateTwo)
	})
	return repositories
}

func formatStringToDateObject(dateString string) time.Time {
	dateTime, _ := time.Parse(time.RFC3339, dateString)
	return dateTime
}

func formatDate(dateString string) string {
	// format to DD/MM/YYYY
	dateTime := formatStringToDateObject(dateString)
	return dateTime.Format(userFacingDateFormat)
}
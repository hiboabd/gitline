package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Owner struct {
	ID int `json:"id"`
}

type Repository struct {
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

type Repositories []Repository

type apiRepositories struct {
	Repository []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Owner     Owner  `json:"owner"`
		URL       string `json:"html_url"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		PushedAt  string `json:"pushed_at"`
		Size      int    `json:"size"`
		Language  string `json:"language"`
	} `json:"repositories"`
}

func GetRepositoryData(c *gin.Context) {
	apiUrl := getEnv("API_URL", "")
	resp, err := http.Get(apiUrl + "/api/v1/repositories")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result apiRepositories
	log.Println(result)
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	userRepositories := formatRepositories(result)

	c.HTML(
		http.StatusOK,
		"timeline.gotmpl",
		gin.H{
			"Repositories": userRepositories,
		},
	)
}

func formatRepositories(r apiRepositories) Repositories {
	var repositories Repositories
	for _, t := range r.Repository {
		var repository = Repository{
			ID:   t.ID,
			Name: t.Name,
			Owner: Owner{
				t.Owner.ID,
			},
			URL:       t.URL,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			PushedAt:  t.PushedAt,
			Size:      t.Size,
			Language:  t.Language,
		}

		repositories = append(repositories, repository)
	}
	return repositories
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
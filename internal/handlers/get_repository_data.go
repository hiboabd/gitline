package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type apiError struct {
	Message string `json:"message"`
}

func (e apiError) Error() string {
	return e.Message
}

type RepositoriesList []apiRepository

type RepositoryData struct {
	Repositories Repositories
}

const userFacingDateFormat string = "02/01/2006"

func (c *Client) GetRepositoryData(username string) (RepositoryData, error) {
	var result RepositoriesList

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/users/%s/repos", username), nil)
	personalAccessToken := GetEnv("PERSONAL_ACCESS_TOKEN", "")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", personalAccessToken))

	if err != nil {
		return RepositoryData{}, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		fmt.Println("error!", err)
		return RepositoryData{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		var notFoundResp apiError
		notFoundResp.Message = "Something went wrong."
		return RepositoryData{}, notFoundResp
	}
	userRepositories := formatRepositories(result)
	return userRepositories, err
}

func formatRepositories(r RepositoriesList) RepositoryData {
	var repositories Repositories
	for _, t := range r {
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
	return RepositoryData{sortedRepositories}
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

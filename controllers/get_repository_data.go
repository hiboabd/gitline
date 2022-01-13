package controllers

import (
	"encoding/json"
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

type RepositoriesList []apiRepository

type apiRepositories struct {
	Repositories []apiRepository `json:"repositories"`
}

type RepositoryData struct {
	Repositories Repositories
}

const userFacingDateFormat string = "02/01/2006"
const timelineTemplate string = "timeline.gotmpl"

func (c *Client) GetRepositoryData() (string, interface{}, error) {
	var result apiRepositories

	req, err := c.newRequest(http.MethodGet, "/api/v1/repositories", nil)
	if err != nil {
		return timelineTemplate, result, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return timelineTemplate, result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	userRepositories := formatRepositories(result)
	return timelineTemplate, userRepositories, err
}

func formatRepositories(r apiRepositories) RepositoryData {
	var repositories Repositories
	for _, t := range r.Repositories {
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

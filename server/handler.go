package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Owner struct {
	ID int `json:"id"`
}

type Repository struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Owner *Owner `json:"owner"`
	URL string `json:"html_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt string `json:"pushed_at"`
	Size int `json:"size"`
	Language string `json:"language"`
}

type Repositories struct {
	Repositories []Repository `json:"repositories"`
}

func GetRepos (w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("response.json")
	if err != nil {
		fmt.Println(err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our repositories array
	var repositories Repositories

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &repositories)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(repositories)
	if err != nil {
		fmt.Println(err)
	}
}

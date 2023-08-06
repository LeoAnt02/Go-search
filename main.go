// Interview: Home Assignment
// Build a Rest API Service in Go to provide a search functionality for this github project:
// https://github.com/avelino/awesome-go
// This API service will have one endpoint:
// “/projects”
// ”/projects” produce a list of all github projects from Readme file from this “awesome-go” project above. See the response payload below
// as expected output.
// “/projects?name={project_name}”, where “project_name” string could be anything like “json” or “audio” or something else.
// The response should be a list of project(s) in json format and the response payload returns a list of project(s) where github project name
// contains {project_name} value.
// For example:
// if we use project_name=”json”, then this project will be in the resulting list:
//
//	{
//	 projects: [
//	 {“url”: “https://github.com/spyzhov/ajson“},
//	 {“url”: “https://github.com/cocoonspace/dynjson“},
//	 ...
//	 ]
//	}

// Assesment: Leo Annette - August 6th 2023
// This is a simple REST API service that provides a search functionality for the github project

//How to run:
// 1. go run main.go
// 2. Open browser and go to http://localhost:8000/projects?name=json
// 3. You will see the result in json format
// 4. You can change the name parameter to any string you want to search for
// 5. You can also go to http://localhost:8000/projects to see all the projects

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type Project struct {
	URL string `json:"url"`
}

var projects []Project

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the 'name' query parameter exists
	key, ok := r.URL.Query()["name"]

	// If 'name' query parameter does not exist, return all projects
	if !ok || len(key[0]) < 1 {
		json.NewEncoder(w).Encode(projects)
		return
	}

	// If 'name' query parameter exists,
	// return projects that contain the 'name' query parameter
	var result []Project
	for _, project := range projects {
		if strings.Contains(project.URL, key[0]) {
			result = append(result, project)
		}
	}

	// If no projects are found, return not found
	if len(result) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func fetchAndParseReadme(gitURL string) {
	// Fetch the README file from the github project
	resp, err := http.Get(gitURL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}
	// Read the response body
	data, _ := ioutil.ReadAll(resp.Body)

	// Find all the github URLs in the README file
	// and store them in the projects slice
	// We can adjust the regex in the future to exclue URLs
	//that are not github project URLs
	re := regexp.MustCompile(`\bhttps://github.com/[a-zA-Z0-9/-]*\b`)
	matches := re.FindAllString(string(data), -1)
	for _, match := range matches {
		projects = append(projects, Project{URL: match})
	}
}

func main() {
	fetchAndParseReadme("https://raw.githubusercontent.com/avelino/awesome-go/master/README.md")

	r := mux.NewRouter()

	r.HandleFunc("/projects", getProjects).Methods("GET")

	http.ListenAndServe(":8000", r)
}

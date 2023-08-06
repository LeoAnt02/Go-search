# Go-search
###### Leo Annette - 6th August 2023
This is a simple REST API service that provides a search functionality for the github project


# GitHub Project Search API in Go

## Introduction

This project is a simple REST API service built in Go that provides search functionality for the GitHub project [Awesome Go](https://github.com/avelino/awesome-go). The API service allows users to search for GitHub projects from the Readme file of the "awesome-go" repository. It has two endpoints, as described below:

1. `/projects`: This endpoint produces a list of all GitHub projects extracted from the Readme file of the "awesome-go" project. The response payload is in JSON format and returns a list of project URLs.

2. `/projects?name={project_name}`: This endpoint takes a query parameter "project_name" as input, which can be any string (e.g., "json" or "audio" or any other term). The response will be a list of projects in JSON format where the GitHub project name contains the specified "project_name" value.


## How to Run

To run the API service locally, follow these steps:

1. Make sure you have Go installed on your machine.

2. Clone this repository to your local machine.

3. Open a terminal or command prompt and navigate to the root directory of the cloned repository.

4. Run the following command to start the API service: 

```` go run main.go ````

5. Once the service is running, open your web browser and go to the following URL to search for projects containing a specific name (e.g., "json"):

http://localhost:8000/projects?name=json


6. The API will return a JSON response containing a list of projects whose names match the search query.

7. To see a list of all projects from the "awesome-go" Readme file, you can visit the following URL:

http://localhost:8000/projects



## Example Response

If you search for projects with the name "json," the API will return a response similar to the following:

```json
{
"projects": [
 { "url": "https://github.com/spyzhov/ajson" },
 { "url": "https://github.com/cocoonspace/dynjson" },
 { ... }
]
}
```

## Note

This API service is built as a home assignment, and it solely serves as an example of creating a basic REST API in Go for searching GitHub projects. It is not intended for production use and might not include advanced error handling or other production-grade features.
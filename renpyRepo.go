package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type APIResponse struct {
	Total_count        int
	Incomplete_results bool
	Items              []Item
}

type GitHubFile struct {
	Name     string
	Path     string
	Sha      string
	Url      string
	Git_url  string
	Html_url string
}

type Item struct {
	GitHubFile
	Repository interface{}
	Score      float64
}

type FileContent struct {
	GitHubFile
	Download_url string
}

func getRenpyFromRepo(repo string) []string {

	// Get all Ren'Py files
	files := getRenpyFilesFromRepo(repo)

	// faster to concatenate big strings with buffers than strings
	var renpyText strings.Builder

	start := time.Now()
	// 4-6s without goroutines
	for _, file := range files {
		fileContentString := getFileContent(file)
		renpyText.WriteString(fileContentString)
	}
	fmt.Println("Fetching files", time.Since(start).Milliseconds())

	totalText := renpyText.String()

	return strings.Split(totalText, "\n")

}

func getFileContent(file Item) string {
	// Get raw text url for a file
	rawFileUrl := strings.Replace(file.Html_url, "github.com", "raw.githubusercontent.com", -1)
	rawFileUrl = strings.Replace(rawFileUrl, "blob/", "", -1)

	resp, err := http.Get(rawFileUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err, resp)
	}
	return string(body)

}

func getRenpyFilesFromRepo(s string) []Item {
	resp, err := http.Get("http://api.github.com/search/code?accept=application/vnd.github.v3+json&q=extension:rpy+repo:" + s)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Convert the body to type string
	sb := string(body)
	// log.Printf(sb)

	a := []byte(sb)

	var apiResponse APIResponse

	if err := json.Unmarshal(a, &apiResponse); err != nil {
		panic(err)
	}

	num := apiResponse.Total_count
	fmt.Println(num)

	return apiResponse.Items
}

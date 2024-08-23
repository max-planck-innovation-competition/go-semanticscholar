package semanticscholar

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/schollz/progressbar/v3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Datasets struct {
	ReleaseID string `json:"release_id"`
	Readme    string `json:"README"`
	Datasets  []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Readme      string `json:"README"`
	} `json:"datasets"`
}

type DownloadLinks struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Readme      string   `json:"README"`
	Files       []string `json:"files"`
}

func MakeRequest(URL string, ApiKey string) (response []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("x-api-key", ApiKey)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
	}
	defer res.Body.Close()

	response, err = io.ReadAll(res.Body)

	return response, err
}

func GetReleaseIds(baseURL string, ApiKey string) (responseData []string, err error) {
	// use faster parser
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Make a GET request to the base URL to get the list of available releases
	response, err := MakeRequest(baseURL, ApiKey)

	var releaseIDs []string
	err = json.Unmarshal(response, &releaseIDs)

	return releaseIDs, err
}

func Download(destPath string, ApiKey string) (err error) {

	var baseURL = "https://api.semanticscholar.org/datasets/v1/release/"
	releaseIDs, err := GetReleaseIds(baseURL, ApiKey)

	for _, releaseID := range releaseIDs {
		// Make a request to get datasets available in the latest release
		datasetsBody, err := MakeRequest(baseURL+releaseID, ApiKey)
		if err != nil {
			log.Fatalf("Failed to make a request for datasets: %v", err)
		}

		var datasets Datasets
		err = json.Unmarshal(datasetsBody, &datasets)
		if err != nil {
			log.Fatalf("Failed to unmarshal datasets JSON: %v", err)
		}

		// Check if the 'abstracts', 'papers' dataset exists
		datasetsList := datasets.Datasets
		papersDatasetExists := false
		abstractsDatasetExists := false
		for _, dataset := range datasetsList {
			if dataset.Name == "papers" {
				papersDatasetExists = true
				print("papersDataset exists")

			}
			if dataset.Name == "abstracts" {
				abstractsDatasetExists = true
				print("abstractsDataset exists")

			}
		}

		if papersDatasetExists {
			log.Println("Start downloading papers-dataset of release:", releaseID)
			// Make a request to get download links for the 'papers' dataset
			datasetName := "papers"
			downloadLinksRequest, err := MakeRequest(baseURL+releaseID+"/dataset/"+datasetName, ApiKey)
			if err != nil {
				log.Fatalf("Failed to create request for download links: %v", err)
			}

			var downloadLinks DownloadLinks
			err = json.Unmarshal(downloadLinksRequest, &downloadLinks)

			for index, link := range downloadLinks.Files {
				newPath := filepath.Join(destPath, releaseID)
				err = os.MkdirAll(newPath, os.ModePerm)
				filePath := filepath.Join(newPath, datasetName+strconv.Itoa(index)+".gz")
				log.Println("Start downloading file:", datasetName+strconv.Itoa(index)+".gz")
				err = downloadFile(link, filePath)
			}
		}

		if abstractsDatasetExists {
			log.Println("Start downloading abstracts-dataset of release:", releaseID)
			// Make a request to get download links for the 'papers' dataset
			datasetName := "abstracts"
			downloadLinksRequest, err := MakeRequest(baseURL+releaseID+"/dataset/"+datasetName, ApiKey)
			if err != nil {
				log.Fatalf("Failed to create request for download links: %v", err)
			}

			var downloadLinks DownloadLinks
			err = json.Unmarshal(downloadLinksRequest, &downloadLinks)

			for index, link := range downloadLinks.Files {
				newPath := filepath.Join(destPath, releaseID)
				err = os.MkdirAll(newPath, os.ModePerm)
				filePath := filepath.Join(newPath, datasetName+strconv.Itoa(index)+".gz")
				log.Println("Start downloading file:", datasetName+strconv.Itoa(index)+".gz")
				err = downloadFile(link, filePath)
			}
		}
	}

	return err
}

// downloadFile downloads a file from the specified URL and saves it to the given filePath
func downloadFile(url, filePath string) error {

	// Get the data from the URL
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
	}
	defer res.Body.Close()

	// Check if the HTTP request was successful
	if res.StatusCode != http.StatusOK {
		return err
	}

	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		res.ContentLength,
		"downloading",
	)

	// Write the data to the file
	_, err = io.Copy(io.MultiWriter(f, bar), res.Body)
	if err != nil {
		return err
	}

	return err
}

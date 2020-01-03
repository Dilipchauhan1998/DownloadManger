package downloader

import (
	"os"
	"strings"
)

//CreateDirectory creates a directory at specified location
func createDirectory(pathToDirectory string, directoryName string) error {
	fullPath := pathToDirectory + "/" + directoryName
	err := os.MkdirAll(fullPath, os.ModePerm)
	return err
}

//FetchFileNameUrl return the fileName of the Url
func fetchFileNameFromUrl(url string) string {
	s := strings.Split(url, "/")
	return s[len(s)-1]
}

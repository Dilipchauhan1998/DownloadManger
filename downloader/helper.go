package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

//SerialDownload downloads the files serially
func serialDownload(currStatus *DownloadStatus, urls []string) {
	err := createDirectory(GLOBAL_PATH, currStatus.Id)
	if err != nil {
		log.Fatal(err)
		currStatus.setStatus(FAILED)
	}

	insertIntoRepository(currStatus)

	for i := range urls {
		fileName := fetchFileNameFromUrl(urls[i])
		filePath := GLOBAL_PATH + currStatus.Id + "/" + fileName

		err := downloadFileFromUrl(filePath, urls[i])
		if err != nil {
			log.Fatal(err)
			currStatus.setStatus(FAILED)
			insertIntoRepository(currStatus)
		}

		currStatus.setFile(filePath, urls[i])
		insertIntoRepository(currStatus)
	}
	currStatus.downloadCompleted()
	insertIntoRepository(currStatus)

}

//ConcurrentDownload downloads the files concurrently
func concurrentDownload(currStatus *DownloadStatus, urls []string) {
	err := createDirectory(GLOBAL_PATH, currStatus.Id)
	if err != nil {
		log.Fatal(err)
		currStatus.setStatus(FAILED)

	}

	insertIntoRepository(currStatus)

	var wg sync.WaitGroup

	for i := range urls {
		fileName := fetchFileNameFromUrl(urls[i])
		filePath := GLOBAL_PATH + currStatus.Id + "/" + fileName

		wg.Add(1)
		go concurrentDownloadHelper(currStatus, filePath, urls[i], &wg)

	}
	wg.Wait()

	currStatus.downloadCompleted()
	insertIntoRepository(currStatus)

}

//ConcurrentDownloadHelper downloads a file and remove a request from wait group
func concurrentDownloadHelper(currStatus *DownloadStatus, filepath string, url string, wg *sync.WaitGroup) {
	err := downloadFileFromUrl(filepath, url)

	if err != nil {
		log.Fatal(err)
		currStatus.setStatus(FAILED)
	}

	currStatus.setFile(filepath, url)
	insertIntoRepository(currStatus)

	wg.Done()

}

//DownloadFileFromUrl download a file at given location
func downloadFileFromUrl(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}


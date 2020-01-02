package main

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/Dilipchauhan1998/DownloadManager/downloadInfo"
	"github.com/Dilipchauhan1998/DownloadManager/utility"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	_ "strconv"
	_ "strings"
	"sync"
	"time"
)

//HealthChecker respond with OK status when Application is up and running
func HealthChecker(c *gin.Context) {
	c.String(200, "OK")
}

//DownloadHandler process the download request based on downLoadType
func DownloadHandler(c *gin.Context) {
	var downloadRequest downloadInfo.DownloadRequest
	c.BindJSON(&downloadRequest)

	//Random id for the current download request
	id := utility.IdGenerator()
	currFileDownloadStatus := downloadInfo.DownloadStatus{Id: id}
	downloadStatus = append(downloadStatus, currFileDownloadStatus)
	downloadStatus[len(downloadStatus)-1].Files = make(map[string]string)
	IndexOfId[id] = len(downloadStatus) - 1

	switch downloadRequest.Type {
	case "serial":
		downloadStatus[IndexOfId[id]].DownloadType = "SERIAL"
		SerialDownloadHandler(downloadRequest.Urls, id)
		c.JSON(http.StatusOK, gin.H{"id": id})

	case "concurrent":
		downloadStatus[IndexOfId[id]].DownloadType = "CONCURRENT"
		go ConcurrentDownloadHandler(downloadRequest.Urls, id)
		c.JSON(http.StatusOK, gin.H{"id": id})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"internal_code": 4001, "message": "unknown type of download"})
	}
}

//SerialDownloadHandler process  the serial download Requests
func SerialDownloadHandler(urls []string, id string) {
	utility.CreateDirectory(downloadLocation, id)
	currFileDownloadLocation := downloadLocation + id

	downloadStatus[IndexOfId[id]].StartTime = time.Now()
	for counter := range urls {
		//fmt.Println("url:",counter, strconv.Itoa(counter))
		downloadStatus[IndexOfId[id]].Status = "QUEUED"
		fileName := utility.FetchFileNameFromUrl(urls[counter])

		DownloadUrlSerial(urls[counter], currFileDownloadLocation+"/"+fileName, id)
	}

	downloadStatus[IndexOfId[id]].EndTime = time.Now()
	downloadStatus[IndexOfId[id]].Status = "SUCCESSFUL"

}

//download the specified url to filepath
func DownloadUrlSerial(url string, filepath string, id string) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	downloadStatus[IndexOfId[id]].Files[url] = filepath
	locationOfFile := downloadInfo.DownloadList{Location: filepath, FileName: utility.FetchFileNameFromUrl(url)}
	downloadList = append(downloadList, locationOfFile)
}

//concurrentDownloadHandler process the concurrent download Requests
func ConcurrentDownloadHandler(urls []string, id string) {
	utility.CreateDirectory(downloadLocation, id)
	currFileDownloadLocation := downloadLocation + id

	downloadStatus[IndexOfId[id]].StartTime = time.Now()
	var wg sync.WaitGroup
	for counter := range urls {
		//fmt.Println("url:",counter, strconv.Itoa(counter))
		downloadStatus[IndexOfId[id]].Status = "QUEUED"
		wg.Add(1)
		fileName := utility.FetchFileNameFromUrl(urls[counter])
		go DownloadUrlConcurrent(urls[counter], currFileDownloadLocation+"/"+fileName, id, &wg)
	}
	wg.Wait()
	downloadStatus[IndexOfId[id]].EndTime = time.Now()
	downloadStatus[IndexOfId[id]].Status = "SUCCESSFUL"

}

//download the specified url to filepath
func DownloadUrlConcurrent(url string, filepath string, id string, wg *sync.WaitGroup) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	downloadStatus[IndexOfId[id]].Files[url] = filepath
	locationOfFile := downloadInfo.DownloadList{Location: filepath, FileName: utility.FetchFileNameFromUrl(url)}
	downloadList = append(downloadList, locationOfFile)
	wg.Done()
}

//Respond to the statusCheck  request of a download
func StatusChecker(c *gin.Context) {
	downloadId := string(c.Param("download_id"))

	if i, ok := IndexOfId[downloadId]; ok {
		files, err := json.Marshal(downloadStatus[i].Files)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		jsonfiles := string(files)
		c.JSON(http.StatusOK, gin.H{"id": downloadId, "start_time": downloadStatus[i].StartTime, "end_time": downloadStatus[i].EndTime,
			"status": downloadStatus[i].Status, "download_type": downloadStatus[i].DownloadType, "files": jsonfiles})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"internal_code": 4002, "message": "unknown download ID"})

	}
}

//ListDownloads respond with listing links to all the downloaded files
func ListDownloads(c *gin.Context) {
	//fmt.Println("downloadLoaction",downloadList)
	c.HTML(http.StatusOK, "downloadedFileList.html", gin.H{"title": "DownloadList", "downloadList": downloadList})
}

func ListFile(c *gin.Context){
	downloadId := string(c.Param("download_id"))
	fileName :=string(c.Param("file_name"))
	fileLocation := downloadLocation+downloadId+"/"+fileName
    fileExtension :=utility.FindExtensionOfFile(fileName)
	c.HTML(http.StatusOK, "showFile.html", gin.H{"title": "downloadedFile", "fileLocation":fileLocation,"extension":fileExtension })
}

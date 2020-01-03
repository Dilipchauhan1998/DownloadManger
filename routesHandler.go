package main

import (
	"github.com/Dilipchauhan1998/DownloadManager/downloader"
	"github.com/gin-gonic/gin"
	"net/http"
)

//DownloadRequest holds the request info of a download
type DownloadRequest struct {
	DownloadType string   `json:"type""`
	Urls         []string `json:"urls""`
}

var request = new(DownloadRequest)

//HealthChecker respond with OK status when Application is up and running
func HealthChecker(c *gin.Context) {
	c.String(200, "OK")
}

//DownloadHandler process the download request
func DownloadHandler(c *gin.Context) {
	c.BindJSON(&request)

	if len(request.Urls) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"internal_code": 4003, "message": "Files list is empty"})
		return
	}

	id := downloader.ProcessDownloadRequest(request.DownloadType, request.Urls)

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"internal_code": 4001, "message": "unknown type of download"})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

//Respond to the statusCheck  request of a download
func StatusChecker(c *gin.Context) {
	downloadId := string(c.Param("downloadId"))

	status := downloader.ProcessStatusRequest(downloadId)

	if status != nil {
		c.JSON(http.StatusOK, gin.H{
			"id":            downloadId,
			"start_time":    status.StartTime,
			"end_time":      status.EndTime,
			"status":        status.Status,
			"download_type": status.DownloadType,
			"files":         status.Files,})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"internal_code": 4002, "message": "unknown download ID"})

	}

}

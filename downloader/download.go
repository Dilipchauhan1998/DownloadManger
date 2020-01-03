package downloader

import (
	_ "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "net/http"
	"time"
)

var SERIAL = "serial"
var CONCURRENT = "concurrent"
var GLOBAL_PATH = "/Users/dilipchauhan/Documents/downloaded_files/"
var SUCCESSFUL = "successful"
var QUEUED = "queued"
var FAILED = "failed"

//DownloadStatus holds the status of a download
type DownloadStatus struct {
	Id           string            `json:"id"`
	StartTime    time.Time         `json:"startTime"`
	EndTime      time.Time         `json:"endTime"`
	Status       string            `json:"status"`
	DownloadType string            `json:"downloadType"`
	Files        map[string]string `json:"files"`
}

//New  initialises the DownloadStatus
func New(downloadType string) DownloadStatus {
	id := uuid.New().String()
	startTime := time.Now()
	status := QUEUED

	ds := DownloadStatus{Id: id, StartTime: startTime, Status: status, DownloadType: downloadType}
	ds.Files = make(map[string]string)
	return ds
}

//setStatus set Status Filed to specified value
func (d *DownloadStatus) setStatus(status string) {
	d.Status = status
}

//setStatus set Files Filed to specified value
func (d *DownloadStatus) setFile(location string, file string) {
	d.Files[file] = location
}

//downloadCompleted set Required Fields after download is completed
func (d *DownloadStatus) downloadCompleted() {
	d.Status = SUCCESSFUL
	d.EndTime = time.Now()
}

//ProcessDownloadRequest process the download request based on download type
func ProcessDownloadRequest(downloadType string, urls []string) string {
	currStatus := New(downloadType)

	switch downloadType {
	case SERIAL:
		go serialDownload(&currStatus, urls)
		return currStatus.Id

	case CONCURRENT:
		go concurrentDownload(&currStatus, urls)
		return currStatus.Id

	default:
		return ""
	}
}

//ProcessStatusRequest returns the status of the download request
func ProcessStatusRequest(id string) *DownloadStatus {
	return fetchFromRepository(id)

}

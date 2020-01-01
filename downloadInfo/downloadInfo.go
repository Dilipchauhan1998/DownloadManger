package downloadInfo

import "time"

//DownloadRequest holds the request info of a download
type DownloadRequest struct{
	Type string `json:"type""`
	Urls []string `json:"urls""`
}

//DownloadStatus holds the current status of the requested download
type DownloadStatus struct{
	Id string `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
	Status string `json:"status"`
	DownloadType string `json:"downloadType"`
	Files map[string]string `json:"files"`
}

//DownloadList holds the Location and Filename of downloaded file
type DownloadList struct{
	Location string `json:"location"`
	FileName string `json:"fileName"`
}


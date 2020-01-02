package main

import (
	"github.com/Dilipchauhan1998/DownloadManager/downloadInfo"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var LengthOfId int = 20
var IndexOfId = make(map[string]int)
var downloadStatus []downloadInfo.DownloadStatus
var downloadList []downloadInfo.DownloadList
var downloadLocation string = "/Users/dilipchauhan/Documents/downloaded_files/"

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	InitializeRoutes()

	router.Run(":8081")
}

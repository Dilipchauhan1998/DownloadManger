package main

//InitializeRoutes redirects to Corresponding handlers of REST endpoint
func InitializeRoutes() {
	router.GET("/health", HealthChecker)

	downloadRoutes := router.Group("/downloads")
	{
		downloadRoutes.POST("/", DownloadHandler)

		downloadRoutes.GET("/:download_id", StatusChecker)
	}
	router.GET("/files", ListDownloads)
	router.GET("/Users/dilipchauhan/Documents/downloaded_files/:download_id/:file_name",ListFile)
}

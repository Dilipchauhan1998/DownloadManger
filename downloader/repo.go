package downloader

//downloads holds the status of all download requests
var downloads = make(map[string]*DownloadStatus)

// InsertIntoRepository insert a download status into repository
func insertIntoRepository(currDownload *DownloadStatus) {
	downloads[currDownload.Id] = currDownload
	//fmt.Println("downloads",downloads)
}

// FetchFromRepository fetch the status of a download  from repository
func fetchFromRepository(id string) *DownloadStatus {

	if _, ok := downloads[id]; ok {
		return downloads[id]
	}
	return nil
}

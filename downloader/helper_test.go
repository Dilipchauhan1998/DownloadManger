package downloader

import (
	"testing"
)

func TestDownloadFileFromUrl(t *testing.T) {
	type DownloadInfo struct {
		filepath string
		url      string
	}

	download := []DownloadInfo{{filepath: GLOBAL_PATH + "Fronalpstock_big.jpg", url: "https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg"},
		{filepath: GLOBAL_PATH + "Fronalpstock_big.jpg", url: "https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG"},
		{filepath: GLOBAL_PATH + "sample-link_1.pdf", url: "https://www.antennahouse.com/XSLsample/pdf/sample-link_1.pdf"},}
	for i := range download {
		err := downloadFileFromUrl(download[i].filepath, download[i].url)
		if err != nil {
			t.Fail()
		}
	}
}

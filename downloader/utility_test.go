package downloader

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	type Directory struct {
		parentDir string
		childDir  string
	}
	childDir1 := uuid.New().String()
	childDir2 := uuid.New().String()
	childDir3 := uuid.New().String()

	directories := []Directory{{parentDir: GLOBAL_PATH, childDir: childDir1},
		{parentDir: GLOBAL_PATH, childDir: childDir2},
		{parentDir: GLOBAL_PATH, childDir: childDir3},}

	for i := range directories {
		err := createDirectory(directories[i].parentDir, directories[i].childDir)
		if err != nil {
			t.Fail()
		}
	}

}

func TestFetchFileNameFromUrl(t *testing.T) {
	type UrlCollection struct {
		url  string
		file string
	}
	urls := []UrlCollection{{url: "https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg", file: "Fronalpstock_big.jpg"},
		{url: "https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG", file: "Big_%26_Small_Pumkins.JPG"},
		{url: "https://www.antennahouse.com/XSLsample/pdf/sample-link_1.pdf", file: "sample-link_1.pdf"},}

	for i := range urls {
		if file := fetchFileNameFromUrl(urls[i].url); file != urls[i].file {
			t.Fail()
		}
	}
}

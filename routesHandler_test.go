package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var CONCURRENT = "concurrent"
var SERIAL = "serial"

type StatusRequest struct {
	Id         string
	StatusCode int
}

type ID struct {
	Id string `json:"id"`
}

//getCollectionOfRequest() returns a collection of download request
func getCollectionOfDownloadRequest() []DownloadRequest {
	requests := []DownloadRequest{{DownloadType: CONCURRENT, Urls: []string{"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG", "https://www.antennahouse.com/XSLsample/pdf/sample-link_1.pdf"}},
		{DownloadType: SERIAL, Urls: []string{"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
			"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG", "https://www.antennahouse.com/XSLsample/pdf/sample-link_1.pdf"}},
		{DownloadType: "no type", Urls: []string{"https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg",
			"https://upload.wikimedia.org/wikipedia/commons/d/dd/Big_%26_Small_Pumkins.JPG", "https://www.antennahouse.com/XSLsample/pdf/sample-link_1.pdf"}},
		{DownloadType: SERIAL, Urls: []string{}},
	}
	return requests
}

func TestHealthChecker(t *testing.T) {
	r := getRouter()

	r.GET("/health", HealthChecker)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/health", nil)

	w := processHttpRequest(r, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}

func TestDownloadHandler(t *testing.T) {
	r := getRouter()

	r.POST("/downloads", DownloadHandler)

	downloadRequests := getCollectionOfDownloadRequest()

	for i := range downloadRequests{
		jsonRequest, err := json.Marshal(downloadRequests[i])
		if err != nil {
			log.Fatal(err)
		}

		req, _ := http.NewRequest("POST", "/downloads", bytes.NewBuffer(jsonRequest))

		w := processHttpRequest(r, req)

		testDownloadResponse(t, downloadRequests[i], w)
	}
}

//testDownloadResponse test whether Response Status match with desired Status
func testDownloadResponse(t *testing.T, req DownloadRequest, res *httptest.ResponseRecorder) {

	switch request.DownloadType {
	case SERIAL:
		if (len(req.Urls) == 0 && res.Code != http.StatusBadRequest) || (len(req.Urls) != 0 && res.Code != http.StatusOK) {
			t.Fail()

		}
	case CONCURRENT:
		if (len(req.Urls) == 0 && res.Code != http.StatusBadRequest) || (len(req.Urls) != 0 && res.Code != http.StatusOK) {
			t.Fail()
		}
	default:
		if res.Code != http.StatusBadRequest {
			t.Fail()
		}
	}
}


func TestStatusChecker(t *testing.T) {

	var testIds  []StatusRequest
	r := getRouter()

	r.POST("/downloads", DownloadHandler)

	downloadRequests := getCollectionOfDownloadRequest()

	for i := range downloadRequests {
		jsonRequest, err := json.Marshal(downloadRequests[i])
		if err != nil {
			log.Fatal(err)
		}

		req, _ := http.NewRequest("POST", "/downloads", bytes.NewBuffer(jsonRequest))

		w := processHttpRequest(r, req)

		if w.Code == http.StatusOK {
			Id := new(ID)
			json.Unmarshal(w.Body.Bytes(), Id)
			testIds = append(testIds, StatusRequest{Id:Id.Id,StatusCode:http.StatusOK})
		}

	}
	//append ids which is not in the repository
	testIds = append(testIds, StatusRequest{Id:uuid.New().String(),StatusCode:http.StatusBadRequest})

	r.GET("/downloads/:downloadId", StatusChecker)
	time.Sleep(1 * time.Second)

	for i := range testIds {
		req, _ := http.NewRequest("GET", "/downloads/"+testIds[i].Id, nil)
		w := processHttpRequest(r, req)

		if w.Code != testIds[i].StatusCode {
			t.Fail()
		}
	}
}

package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"vollkorntomate.de/cvedash-api/internal/config"
	"vollkorntomate.de/cvedash-api/internal/tools"
)

type NVDRequestParams struct {
	LastModStartDate string
	LastModEndDate   string
	StartIndex       int
}

var HTTPClient http.Client

func ImportNewEntries() {
	reqParams := NVDRequestParams{}
	lastUpdate := DB.GetLastNVDUpdate()
	printProgress := false
	log.Println("Starting import of new entries. Last update was", lastUpdate.LastUpdate)
	currentTime := time.Now()
	if !lastUpdate.LastUpdate.IsZero() && lastUpdate.Version == DB_VERSION {
		reqParams.LastModStartDate = tools.FormatISODate(lastUpdate.LastUpdate)
		reqParams.LastModEndDate = tools.FormatISODate(currentTime)
	} else {
		printProgress = true
	}

	var err error

	resultsPerPage := 0
	total := 1
	for offset := 0; offset < total; offset += resultsPerPage {
		reqParams.StartIndex = offset
		var response *NVDResponse
		response, err = makeNVDRequest(reqParams)
		if err != nil {
			log.Println("Import of new entries failed:", err)
			return
		}
		if offset == 0 {
			resultsPerPage = response.ResultsPerPage
			total = response.TotalResults
			log.Println("Total number of new entries:", total)
		}
		if printProgress {
			fmt.Printf("\rProgress: %d of %d (%.1f%%)", offset, total, (float32(offset)/float32(total))*100.0)
		}

		minimal := response.ConvertToMinimal()
		DB.SaveBatch(minimal)

		apiSleep()
	}
	if printProgress {
		fmt.Print("\n")
	}

	if err == nil {
		DB.UpdateLastNVDUpdate(currentTime)
		log.Println("Import finished successfully")
	} else {
		log.Println("Import failed", err)
	}
}

func makeNVDRequest(params NVDRequestParams) (*NVDResponse, error) {
	baseURL := "https://services.nvd.nist.gov/rest/json/cves/2.0"

	query := url.Values{}
	query.Set("startIndex", strconv.Itoa(params.StartIndex))

	if params.LastModStartDate != "" && params.LastModEndDate != "" {
		query.Set("lastModStartDate", params.LastModStartDate)
		query.Set("lastModEndDate", params.LastModEndDate)
	}

	url := baseURL + "?" + query.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Could not create HTTP request", err)
		return nil, err
	}
	setAPIKeyHeader(req)
	response, err := HTTPClient.Do(req)
	if err != nil {
		log.Println("Could not get CVEs from NVD", err)
		return nil, err
	}
	defer response.Body.Close()

	ret := NVDResponse{}
	jsonParser := json.NewDecoder(response.Body)
	err = jsonParser.Decode(&ret)
	if err != nil {
		log.Println("Could not decode NVD response JSON", err)
		return nil, err
	}

	return &ret, nil
}

func setAPIKeyHeader(req *http.Request) {
	if config.Config.NVDAPIKey != "" {
		req.Header.Set("apiKey", config.Config.NVDAPIKey)
	}
}

func apiSleep() {
	if config.Config.NVDAPIKey != "" {
		time.Sleep(1 * time.Second)
	} else {
		time.Sleep(6 * time.Second)
	}
}

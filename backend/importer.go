package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type NVDRequestParams struct {
	LastModStartDate string
	LastModEndDate   string
	StartIndex       int
}

func ImportNewEntries() {
	reqParams := NVDRequestParams{}
	lastUpdate := DB.GetLastNVDUpdate()
	if !lastUpdate.IsZero() {
		reqParams.LastModStartDate = FormatISODate(lastUpdate)
		reqParams.LastModEndDate = FormatISODate(time.Now())
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
		}
		resultsPerPage = response.ResultsPerPage
		total = response.TotalResults

		minimal := make([]MinimalCVEData, resultsPerPage)
		for i := 0; i < resultsPerPage; i++ {
			minimal[i] = response.Vulnerabilities[i].CVE.ToMinimalCVEData()
		}

		DB.SaveBatch(minimal)

		time.Sleep(6 * time.Second)
	}

	if err == nil {
		DB.UpdateLastNVDUpdate()
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
	response, err := http.Get(url)
	if err != nil {
		log.Println("Could not get CVEs from NVD", err)
		return nil, err
	}

	ret := NVDResponse{}
	jsonParser := json.NewDecoder(response.Body)
	err = jsonParser.Decode(&ret)
	if err != nil {
		log.Println("Could not decode NVD response JSON", err)
		return nil, err
	}

	return &ret, nil
}

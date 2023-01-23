package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const apiResultsPerPage = 10

var statsCache map[string]CVEStats

func RunAPIServer() {
	statsCache = make(map[string]CVEStats)
	updateCache()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/latest/{page:[0-9]+}", getLatestPublishedCVEs)
	router.Get("/stats/{duration:24h|7d|30d|1y|ytd}", getStats)

	log.Println("Starting API server")
	http.ListenAndServe(":3000", router)
}

func getLatestPublishedCVEs(response http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(chi.URLParam(req, "page"))
	offset := (page - 1) * apiResultsPerPage

	var cves []MinimalCVEData
	DB.DB.Order("published DESC").Limit(apiResultsPerPage).Offset(offset).Find(&cves)

	writeJSONResponse(response, cves)
}

func getStats(response http.ResponseWriter, req *http.Request) {
	duration := chi.URLParam(req, "duration")

	stats, ok := statsCache[duration]
	if !ok {
		response.WriteHeader(404)
		return
	}

	writeJSONResponse(response, stats)
}

func writeJSONResponse(response http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Could not marshal to JSON")
		response.WriteHeader(500)
		return
	}

	response.Write(jsonData)
}

func updateCache() {
	now := time.Now().UTC()
	timeAgo24h := now.Add(-24 * time.Hour)
	timeAgo7d := now.Add(-7 * 24 * time.Hour)
	timeAgo30d := now.Add(-30 * 24 * time.Hour)
	timeAgo1y := time.Date(now.Year()-1, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
	timeAgoYTD := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	statsCache["24h"] = DB.GetStats(timeAgo24h)
	statsCache["7d"] = DB.GetStats(timeAgo7d)
	statsCache["30d"] = DB.GetStats(timeAgo30d)
	statsCache["1y"] = DB.GetStats(timeAgo1y)
	statsCache["ytd"] = DB.GetStats(timeAgoYTD)
}

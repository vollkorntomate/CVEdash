package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"vollkorntomate.de/cvedash-api/internal/config"
	"vollkorntomate.de/cvedash-api/internal/data"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
)

const apiResultsPerPage = 10

var StatsCache CVEStatsCache

func RunAPIServer(ctx context.Context) {
	StatsCache = NewStatsCache()
	UpdateStateCache()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(httprate.LimitByIP(int(config.Config.APIRateLimitPerMinute), time.Minute))

	router.Options("/*", corsOptions)
	router.Get("/latest/{page:[0-9]+}", getLatestPublishedCVEs)
	router.Get("/stats", getAllStats)
	router.Get("/stats/{duration:24h|7d|30d|1y|ytd|alltime}", getStats)

	server := &http.Server{Addr: ":8077", Handler: router}
	go func() {
		log.Println("Starting API server on http://localhost:8077")
		server.ListenAndServe()
	}()
	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
	}()
}

func corsOptions(response http.ResponseWriter, req *http.Request) {
	response.Header().Add("Access-Control-Request-Method", http.MethodGet)
	response.Header().Add("Access-Control-Request-Headers", "Content-Type")
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.WriteHeader(http.StatusNoContent)
}

func getLatestPublishedCVEs(response http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(chi.URLParam(req, "page"))
	offset := (page - 1) * apiResultsPerPage

	var cves []data.MinimalCVEData
	data.DB.DB.Where("status <> ?", "Rejected").Order("published DESC").Limit(apiResultsPerPage).Offset(offset).Find(&cves)

	writeJSONResponse(response, cves)
}

func getStats(response http.ResponseWriter, req *http.Request) {
	duration := chi.URLParam(req, "duration")

	stats, ok := StatsCache.Get(duration)
	if !ok {
		response.WriteHeader(404)
		return
	}

	writeJSONResponse(response, stats)
}

func getAllStats(response http.ResponseWriter, req *http.Request) {
	stats := StatsCache.GetAll()

	writeJSONResponse(response, stats)
}

func writeJSONResponse(response http.ResponseWriter, data any) {
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Header().Add("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Could not marshal to JSON")
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Write(jsonData)
}

func UpdateStateCache() {
	now := time.Now().UTC()
	timeAgo24h := now.Add(-24 * time.Hour)
	timeAgo7d := now.Add(-7 * 24 * time.Hour)
	timeAgo30d := now.Add(-30 * 24 * time.Hour)
	timeAgo1y := time.Date(now.Year()-1, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
	timeAgoYTD := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	timeAgoAlltime := time.Time{} // zero-time

	StatsCache.Set("24h", data.DB.GetStats(timeAgo24h))
	StatsCache.Set("7d", data.DB.GetStats(timeAgo7d))
	StatsCache.Set("30d", data.DB.GetStats(timeAgo30d))
	StatsCache.Set("1y", data.DB.GetStats(timeAgo1y))
	StatsCache.Set("ytd", data.DB.GetStats(timeAgoYTD))
	StatsCache.Set("alltime", data.DB.GetStats(timeAgoAlltime))
}

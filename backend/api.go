package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const apiResultsPerPage = 10

func RunAPIServer() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/latest/{page:[0-9]+}", getLatestPublishedCVEs)

	log.Println("Starting API server")
	http.ListenAndServe(":3000", router)
}

func getLatestPublishedCVEs(response http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(chi.URLParam(req, "page"))
	offset := (page - 1) * apiResultsPerPage

	var cves []MinimalCVEData
	DB.DB.Order("published DESC").Limit(apiResultsPerPage).Offset(offset).Find(&cves)

	jsonData, err := json.Marshal(cves)
	if err != nil {
		log.Println("Could not marshal to JSON")
		response.WriteHeader(500)
	}

	response.Write(jsonData)
}

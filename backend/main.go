package main

import "net/http"

const DB_VERSION = 1

var HTTPClient http.Client

func main() {
	Config = LoadConfig("config.json")

	DB = NewDB(Config.DBFilePath)
	DB.InitDB()

	HTTPClient = http.Client{}

	ImportNewEntries()
	RunAPIServer()
}

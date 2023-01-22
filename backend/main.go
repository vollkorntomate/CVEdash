package main

import "net/http"

var HTTPClient http.Client

func main() {
	Config = LoadConfig("config.json")

	DB = NewDB(Config.DBFilePath)
	DB.InitDB()

	HTTPClient = http.Client{}

	ImportNewEntries()
}

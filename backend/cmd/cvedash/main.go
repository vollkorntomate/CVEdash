package main

import (
	"net/http"
	"sync"
	"time"

	"vollkorntomate.de/cvedash/internal/api"
	"vollkorntomate.de/cvedash/internal/config"
	"vollkorntomate.de/cvedash/internal/data"
)

func main() {
	data.HTTPClient = http.Client{}
	config.Config = config.LoadConfig("config.json")

	data.DB = data.NewDB(config.Config.DBFilePath)
	data.DB.InitDB()

	data.ImportNewEntries()
	api.RunAPIServer()

	ScheduleBackgroundTasks()
}

func ScheduleBackgroundTasks() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(time.Minute * time.Duration(config.Config.MinutesBetweenNVDUpdates)) {
			data.ImportNewEntries()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(time.Minute * 5) {
			api.UpdateStateCache()
		}
	}()

	wg.Wait()
}

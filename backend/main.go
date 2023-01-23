package main

import (
	"net/http"
	"sync"
	"time"
)

const DB_VERSION = 1

var HTTPClient http.Client

func main() {
	Config = LoadConfig("config.json")

	DB = NewDB(Config.DBFilePath)
	DB.InitDB()

	HTTPClient = http.Client{}

	ImportNewEntries()
	RunAPIServer()

	ScheduleBackgroundTasks()
}

func ScheduleBackgroundTasks() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(time.Minute * time.Duration(Config.MinutesBetweenUpdates)) {
			ImportNewEntries()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(time.Minute * 5) {
			UpdateStateCache()
		}
	}()

	wg.Wait()
}

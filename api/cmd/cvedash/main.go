package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"vollkorntomate.de/cvedash-api/internal/api"
	"vollkorntomate.de/cvedash-api/internal/config"
	"vollkorntomate.de/cvedash-api/internal/data"
)

func main() {
	data.HTTPClient = http.Client{}
	config.Config = config.LoadConfig("config.json")

	data.DB = data.NewDB(config.Config.DBFilePath)
	data.DB.InitDB()

	data.ImportNewEntries()
	api.RunAPIServer()

	MakeInterruptHandler()

	ScheduleBackgroundTasks()
}

func MakeInterruptHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Bye!")
		os.Exit(0)
	}()
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

package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())

	data.HTTPClient = http.Client{}
	config.Config = config.LoadConfig("config.json")

	data.DB = data.NewDB(config.Config.DBFilePath)
	data.DB.InitDB()

	data.ImportNewEntries()
	api.RunAPIServer(ctx)

	MakeInterruptHandler(cancel)

	ScheduleBackgroundTasks(ctx)

	log.Println("Bye!")
}

func MakeInterruptHandler(cancel context.CancelFunc) {
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

		<-exit // wait for interrupt
		log.Println("Shutting down...")
		cancel()
	}()
}

func ScheduleBackgroundTasks(ctx context.Context) {
	wg := sync.WaitGroup{}

	importTicker := time.NewTicker(time.Minute * time.Duration(config.Config.MinutesBetweenNVDUpdates))
	defer importTicker.Stop()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-importTicker.C:
				data.ImportNewEntries()
			case <-ctx.Done():
				return
			}
		}
	}()

	cacheUpdateTicker := time.NewTicker(time.Minute * 5)
	defer cacheUpdateTicker.Stop()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-cacheUpdateTicker.C:
				api.UpdateStateCache()

			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
}

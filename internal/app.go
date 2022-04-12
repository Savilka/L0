package internal

import (
	"L0/internal/api"
	"L0/internal/caching"
	"L0/internal/store"
	"L0/internal/streaming"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// App is main structure of service
type App struct {
	stream streaming.Streaming
	api    api.Api
}

// Run running streaming service and server
func (a *App) Run(stanClusterID, clientID, URL string) {
	err := a.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = a.stream.ConnectAndSubscribe(stanClusterID, clientID, URL)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8080", a.api.Router)
	if err != nil {
		log.Fatal(err)
	}
}

// Init initialize api, cache, store and streaming service
func (a *App) Init() error {
	cache := caching.NewCache()

	if err := godotenv.Load("./conf/.env"); err != nil {
		log.Print("No .env file found")
	}

	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return errors.New("DATABASE_URL not exist")
	}
	newStore, err := store.NewStore(databaseURL)
	if err != nil {
		return err
	}

	a.api.InitRouter(cache)

	a.stream.InitStreaming(cache, newStore)

	err = newStore.RestoreCache(cache)
	if err != nil {
		return err
	}

	return nil
}

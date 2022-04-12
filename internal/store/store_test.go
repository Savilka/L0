package store

import (
	"L0/internal/caching"
	"L0/internal/model"
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var store *Store

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../conf/.env"); err != nil {
		log.Print("No .env file found")
		os.Exit(1)
	}

	databaseURL, exists := os.LookupEnv("DATABASE_URL_TEST")
	if !exists {
		log.Println("DATABASE_URL not exist")
		os.Exit(1)
	}

	var err error
	store, err = NewStore(databaseURL)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func clearTable() {
	_, err := store.dbpool.Exec(context.Background(), "truncate table orders  restart identity")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func TestBadUrl(t *testing.T) {
	url := "badURL"
	store, err := NewStore(url)
	assert.Equal(t, err.Error(), "cannot parse `badURL`: failed to parse as DSN (invalid dsn)")
	assert.Equal(t, store, Store{})
}

func TestStore_AddOrder(t *testing.T) {
	err := store.AddOrder(model.Order{})
	assert.Equal(t, err, nil)
	clearTable()
}

func TestStore_GetOrder(t *testing.T) {
	err := store.AddOrder(model.Order{OrderUID: "1"})
	assert.Equal(t, err, nil)

	order, err := store.getOrder("1")
	assert.Equal(t, order, model.Order{OrderUID: "1"})
	assert.Equal(t, err, nil)

	order, err = store.getOrder("2")
	assert.Equal(t, order, model.Order{})
	assert.Equal(t, err.Error(), "no rows in result set")
	clearTable()
}

func TestStore_RestoreCash(t *testing.T) {
	err := store.AddOrder(model.Order{OrderUID: "1"})
	assert.Equal(t, err, nil)
	err = store.AddOrder(model.Order{OrderUID: "2"})
	assert.Equal(t, err, nil)

	cache := caching.NewCache()
	err = store.RestoreCache(cache)
	assert.Equal(t, err, nil)
	order, err := cache.Get("1")
	assert.Equal(t, err, nil)
	assert.Equal(t, order, model.Order{OrderUID: "1"})
	order, err = cache.Get("2")
	assert.Equal(t, err, nil)
	assert.Equal(t, order, model.Order{OrderUID: "2"})

	clearTable()
}

package caching

import (
	"L0/internal/model"
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

var cache *Cache

func TestMain(m *testing.M) {
	cache = NewCache()
	code := m.Run()
	os.Exit(code)
}

func TestCache(t *testing.T) {
	cache.Put(model.Order{OrderUID: "1"})
	assert.Equal(t, true, cache.IsExist("1"))
	order, err := cache.Get("1")
	assert.Equal(t, err, nil)
	assert.Equal(t, model.Order{OrderUID: "1"}, order)

	order, err = cache.Get("2")
	assert.Equal(t, err.Error(), "the order isn't in cache")
	assert.Equal(t, model.Order{}, order)
}

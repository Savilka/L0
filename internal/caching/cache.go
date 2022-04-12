package caching

import (
	"L0/internal/model"
	"errors"
	"sync"
)

// Cache is a structure for in memory cashing
type Cache struct {
	data map[string]model.Order
	mu   sync.RWMutex
}

// NewCache return initialized `Cache `struct
func NewCache() *Cache {
	return &Cache{data: make(map[string]model.Order), mu: sync.RWMutex{}}
}

// Get return order from cache by id
func (c *Cache) Get(id string) (model.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	order, ok := c.data[id]
	if !ok {
		return model.Order{}, errors.New("the order isn't in cache")
	}

	return order, nil
}

// Put order in cache
func (c *Cache) Put(order model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[order.OrderUID] = order
}

// IsExist true if id exist in cache
func (c *Cache) IsExist(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[id]

	return ok
}

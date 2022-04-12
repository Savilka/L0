package caching

import (
	"L0/internal/model"
	"errors"
	"sync"
)

type Cache struct {
	data map[string]model.Order
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]model.Order), mu: sync.RWMutex{}}
}

func (c *Cache) Get(id string) (model.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	order, ok := c.data[id]
	if !ok {
		return model.Order{}, errors.New("the order isn't in cache")
	}

	return order, nil
}

func (c *Cache) Put(order model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[order.OrderUID] = order
}

func (c *Cache) IsExist(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[id]

	return ok
}

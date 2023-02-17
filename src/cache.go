package main

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex

	values    map[string]item
	expiresAt time.Duration
}

func NewCache(expiresAt time.Duration) *Cache {
	
	cache := Cache{
		values: make(map[string]item),
		expiresAt: expiresAt,
	}

	if expiresAt > 0 {
		cache.StartGC()
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	_, ok := c.values[key]
	if ok {
		return errors.New("this key is already exists in cache")
	} 

	expiration := time.Now().Add(duration).UnixNano()

	c.Lock()
	defer c.Unlock()

	c.values[key] = CreateNewItem(value, expiration)

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()
	
	item, ok := c.values[key]
	
	if !ok {
		return nil, errors.New("no such key in cache")
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, errors.New("this value is expired")
		}

	}
	return item, nil
	
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	_, ok := c.values[key]
	if ok {
		delete(c.values, key)
		return nil
	}

	return errors.New("no such key in cache")
}
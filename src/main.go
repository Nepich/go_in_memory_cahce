package main

import (
	"errors"
	"fmt"
)


type Cache struct {
	values map[string]interface{}
}

func NewCache() *Cache {
	stor := make(map[string]interface{})

	cache := Cache {stor}

	return &cache
}

func (c *Cache) Set(key string, value interface{}) error {
	_, ok := c.values[key]
	if ok {
		return errors.New("this key is already exists in cache")
	} 
	c.values[key] = value
	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.values[key]
	if ok {
		return value, nil
	}
	return nil, errors.New("no such key in cache")
}

func (c *Cache) Delete(key string) error {
	_, ok := c.values[key]
	if ok {
		delete(c.values, key)
		return nil
	}

	return errors.New("no such key in cache")
}

func main() {
	cache := NewCache()
	fmt.Println(cache)
}
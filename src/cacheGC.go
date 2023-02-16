package main

import "time"

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {

	for {
		<-time.After(c.expiresAt)

		if c.values == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.CleanItems(keys)

		}
	}
}

func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()

	for k, i := range c.values {
		if i.CheckLifeTime() {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) CleanItems(keys []string) {
	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.values, k)
	}
}
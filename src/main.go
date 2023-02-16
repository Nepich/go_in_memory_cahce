package main

import (
	"fmt"
	"time"
)

func main() {
	cache := NewCache(time.Minute * 10)
	cache.Set("1", "New object", time.Minute)

	fmt.Println(cache.Get("1"))
}
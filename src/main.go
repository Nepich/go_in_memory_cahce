package main

import (
	"fmt"
	"time"
)

func main() {
	cache := NewCache(time.Minute * 10)
	cache.Set("1", "New object", time.Second*5)
	fmt.Println(cache.Get("1"))
	time.Sleep(time.Second*10)
	fmt.Println(cache.Get("1"))
}
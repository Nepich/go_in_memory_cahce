package main

import "time"

type item struct {
	Value          interface{}
	CreatedAt time.Time
	Expiration int64
}

func CreateNewItem(value any, expiration int64) item {
	
	return item{
		Value: value,
		CreatedAt: time.Now(),
		Expiration: expiration,
	}
}

//CheckLifeTime returns true if item is expired and false if not
func (i *item) CheckLifeTime() bool {
	
	return time.Unix(0, (i.CreatedAt.UnixNano() + i.Expiration)).After(time.Now())
}


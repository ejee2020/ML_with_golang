package main

import (

	cache "github.com/patrickmn/go-cache"
	"fmt"
	"time"

	
)

func main() {

	c := cache.New(5 * time.Minute, 30 * time.Second)
	c.Set("mykey", "myvalue", cache.DefaultExpiration)
	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}
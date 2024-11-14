package main

import (
	"fmt"
	"time"

	"github.com/TheoBrigitte/expirymap"
)

func main() {
	// Create a new map[int]string with an expiry delay of 5s and a garbage collection interval of 1s.
	m := expirymap.New[int, string](time.Second*5, time.Second)
	defer m.Stop()

	// Set 1=foo, 2=bar, 3=baz in the map.
	m.Set(1, "foo")
	m.Set(2, "bar")
	m.Set(3, "baz")

	// Iterate over the map and print the key and value.
	for k, v := range m.Iterate() {
		fmt.Println(k, v)
	}
}

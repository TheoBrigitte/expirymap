package main

import (
	"fmt"
	"time"

	"github.com/TheoBrigitte/expirymap"
)

func main() {
	// Define a key and a value.
	key := 1
	value := []string{"foo", "bar", "baz"}

	// Create a new expiry map of type map[int][]string
	// with an expiry delay of 1ns and a garbage collection interval of 1ms.
	m := expirymap.New[int, []string](time.Nanosecond, time.Millisecond)
	defer m.Stop()

	// Set 1=[foo bar baz] in the map.
	m.Set(key, value)

	fmt.Println(m.Get(1))            // [foo bar baz]
	time.Sleep(time.Millisecond * 2) // Wait for the entry to expire.
	fmt.Println(m.Get(1))            // []
}

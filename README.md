<p align="center">
  <a href="https://pkg.go.dev/github.com/TheoBrigitte/expirymap"><img src="https://pkg.go.dev/badge/github.com/TheoBrigitte/expirymap.svg" alt="Go Reference"></a>
  <a href="https://github.com/TheoBrigitte/expirymap/actions/workflows/test.yaml"><img src="https://github.com/TheoBrigitte/expirymap/actions/workflows/test.yaml/badge.svg" alt="Github action"></a>
  <a href="http://github.com/TheoBrigitte/expirymap/releases"><img src="https://img.shields.io/github/release/TheoBrigitte/expirymap.svg" alt="Latest Release"></a>
</p>

## ExpiryMap

This Go package provides a map that automatically removes entries after a given expiry delay.

### Features

* The map key can be any comparable type
* The map value can be any type
* The map is safe for concurrent use
* The expiry delay is specified as a `time.Duration` value

### Methods

* `New` - creates a new `Map`
* `Get`, `Set`, `Delete` - standard map operations
* `Len` - returns the number of entries in the map
* `Iterate` - iterates over all entries in the map
* `Clear` - removes all entries from the map
* `Stop` - stops the background goroutine that removes expired entries

### Example

```go
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
```
source [example/simple/simple.go](./example/simple/simple.go)

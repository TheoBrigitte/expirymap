<p align="center">
  <a href="https://pkg.go.dev/github.com/TheoBrigitte/expirymap"><img src="https://pkg.go.dev/badge/github.com/TheoBrigitte/expirymap.svg" alt="Go Reference"></a>
  <a href="https://github.com/TheoBrigitte/expirymap/actions/workflows/test.yml"><img src="https://github.com/TheoBrigitte/expirymap/actions/workflows/test.yml/badge.svg" alt="Github action"></a>
  <a href="github.com/TheoBrigitte/expirymap/releases"><img src="https://img.shields.io/github/release/TheoBrigitte/expirymap.svg" alt="Latest Release"></a>
</p>

## ExpiryMap

This Go package provides a map that automatically removes entries after a given expiry delay.

### Features

* The map key can be any comparable type
* The map value can be any type
* The map is safe for concurrent use
* The expiry delay is specified as a `time.Duration` value

### Methods

* NewExpiryMap - creates a new ExpiryMap
* Get, Set, Delete - standard map operations
* Len - returns the number of entries in the map
* Iterate - iterates over all entries in the map
* Clear - removes all entries from the map
* Stop - stops the background goroutine that removes expired entries

### Example

See [example/simple/simple.go](./example/simple/simple.go)

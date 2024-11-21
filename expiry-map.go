package expirymap

import (
	"iter"
	"sync"
	"time"
)

// Map is a map of K key and V values with a builtin garbage cleaner that automatically
// deletes entries after a given expiry delay.
type Map[K comparable, V any] struct {
	storedMap            map[K]*Content[V]
	expiryDelay          time.Duration
	gargabeCleanInterval time.Duration
	mutex                sync.Mutex
	stop                 chan bool
}

// Content is the value stored for a given key in the ExpiryMap.
type Content[V any] struct {
	Data V

	lastUpdated time.Time
}

// New returns a new ExpiryMap.
// It also starts a goroutine that periodically cleans up expired entries
// according to the expiryDelay every gargabeCleanInterval.
func New[K comparable, V any](expiryDelay, gargabeCleanInterval time.Duration) *Map[K, V] {
	s := &Map[K, V]{
		storedMap:            make(map[K]*Content[V]),
		expiryDelay:          expiryDelay,
		gargabeCleanInterval: gargabeCleanInterval,
		stop:                 make(chan bool),
	}

	s.start()

	return s
}

// Get returns the value for a given key.
func (s *Map[K, V]) Get(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	content, found := s.storedMap[key]

	if !found {
		content = &Content[V]{}
	}

	return content.Data, found
}

// Set sets the value for a given key and reset its expiry time.
func (s *Map[K, V]) Set(key K, data V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	content := &Content[V]{}
	content.lastUpdated = time.Now()
	content.Data = data
	s.storedMap[key] = content
}

// Delete deletes the value for a given key.
func (s *Map[K, V]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.storedMap, key)
}

// Len returns the number of stored entries.
func (s *Map[K, V]) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.storedMap)
}

// Iterate returns an iterator to loop over the stored entries.
func (s *Map[K, V]) Iterate() iter.Seq2[K, V] {
	return func(next func(K, V) bool) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		for k, v := range s.storedMap {
			if !next(k, v.Data) {
				return
			}
		}
	}
}

// Clear deletes all stored entries.
func (s *Map[K, V]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	clear(s.storedMap)
}

// Stop stops the garbage cleaner goroutine.
func (s *Map[K, V]) Stop() {
	s.stop <- true
}

func (s *Map[K, V]) start() {
	go func() {
		for {
			select {
			case <-s.stop:
				return
			case <-time.Tick(s.gargabeCleanInterval):
				s.gargabeClean()
			}
		}
	}()
}

func (s *Map[K, V]) gargabeClean() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	expiredTime := time.Now().Add(-s.expiryDelay)

	for key, u := range s.storedMap {
		if u.lastUpdated.Before(expiredTime) {
			delete(s.storedMap, key)
		}
	}
}

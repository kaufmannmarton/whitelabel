package storage

// Source: https://github.com/goenning/go-cache-demo/blob/master/cache/memory/cache.go

import (
	"sync"
	"time"
)

// Item is a cached reference
type item struct {
	Content    []byte
	Expiration int64
}

// Expired returns true if the item has expired.
func (i item) Expired() bool {
	if i.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > i.Expiration
}

//Storage mechanism for caching strings in memory
type MemoryStorage struct {
	items map[string]item
	mu    *sync.RWMutex
}

//NewStorage creates a new in memory storage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		items: make(map[string]item),
		mu:    &sync.RWMutex{},
	}
}

//Get a cached content by key
func (s MemoryStorage) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item := s.items[key]
	if item.Expired() {
		delete(s.items, key)
		return nil
	}
	return item.Content
}

//Set a cached content by key
func (s MemoryStorage) Set(key string, content []byte, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

package tzcache

import (
	"sync"
	"time"
)

type TzCache struct {
	lock  sync.Mutex
	cache map[string]*time.Location
}

// New creates a new TzCache instance.
func New() *TzCache {
	return &TzCache{
		cache: make(map[string]*time.Location),
	}
}

// Get retrieves the time.Location for the given timezone name from the cache or loads it if not present. Returns error if the timezone cannot be loaded.
func (c *TzCache) Get(name string) (*time.Location, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if loc, ok := c.cache[name]; ok {
		return loc, nil
	}

	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	c.cache[name] = loc

	return loc, nil
}

// MustGet retrieves the time.Location for the given timezone name from the cache or loads it if not present. It panics if the timezone cannot be loaded.
func (c *TzCache) MustGet(name string) *time.Location {
	loc, err := c.Get(name)
	if err != nil {
		panic(err)
	}
	return loc
}

// Clear removes all entries from the cache.
func (c *TzCache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache = make(map[string]*time.Location)
}

// Delete removes a specific timezone from the cache.
func (c *TzCache) Delete(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.cache, name)
}

// Size returns the number of entries in the cache.
func (c *TzCache) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	return len(c.cache)
}

// Contains checks if the cache contains a specific timezone.
func (c *TzCache) Contains(name string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, exists := c.cache[name]

	return exists
}

// List returns a slice of all timezone names currently in the cache.
func (c *TzCache) List() []string {
	c.lock.Lock()
	defer c.lock.Unlock()

	names := make([]string, 0, len(c.cache))
	for name := range c.cache {
		names = append(names, name)
	}

	return names
}

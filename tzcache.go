package tzcache

import (
	"time"
)

// TzCache is an interface for a time zone cache.
type TzCache interface {
	// Location retrieves the time.Location for the given timezone name from the cache or loads it if not present. Returns error if the timezone cannot be loaded.
	Location(name string) (*time.Location, error)

	// MustLocation retrieves the time.Location for the given timezone name from the cache or loads it if not present. It panics if the timezone cannot be loaded.
	MustLocation(name string) *time.Location

	// Delete removes a specific timezone from the cache.
	Delete(name string)

	// Contains checks if the cache contains a specific timezone.
	Contains(name string) bool

	// Size returns the number of entries in the cache.
	Size() int

	// List returns a slice of all timezone names currently in the cache.
	List() []string

	// Clear removes all entries from the cache.
	Clear()
}

// NewSafe creates a new thread safe TzCache instance.
func NewSafe() TzCache {
	return &TzCacheSafe{
		cache: make(map[string]*time.Location),
	}
}

// NewUnsafe creates a new thread unsafe TzCache instance (without locking). This is should only be used in single-threaded contexts.
func NewUnsafe() TzCache {
	return &TzCacheUnsafe{
		cache: make(map[string]*time.Location),
	}
}

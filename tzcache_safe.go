package tzcache

import (
	"sync"
	"time"
)

type TzCacheSafe struct {
	lock  sync.Mutex
	cache map[string]*time.Location
}

func (c *TzCacheSafe) Location(name string) (*time.Location, error) {
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

func (c *TzCacheSafe) MustLocation(name string) *time.Location {
	loc, err := c.Location(name)
	if err != nil {
		panic(err)
	}
	return loc
}

func (c *TzCacheSafe) Delete(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.cache, name)
}

func (c *TzCacheSafe) Contains(name string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, exists := c.cache[name]

	return exists
}

func (c *TzCacheSafe) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	return len(c.cache)
}

func (c *TzCacheSafe) List() []string {
	c.lock.Lock()
	defer c.lock.Unlock()

	names := make([]string, 0, len(c.cache))
	for name := range c.cache {
		names = append(names, name)
	}

	return names
}

func (c *TzCacheSafe) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache = make(map[string]*time.Location)
}

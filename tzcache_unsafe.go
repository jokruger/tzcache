package tzcache

import (
	"time"
)

type TzCacheUnsafe struct {
	cache map[string]*time.Location
}

func (c *TzCacheUnsafe) Location(name string) (*time.Location, error) {
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

func (c *TzCacheUnsafe) MustLocation(name string) *time.Location {
	loc, err := c.Location(name)
	if err != nil {
		panic(err)
	}
	return loc
}

func (c *TzCacheUnsafe) Delete(name string) {
	delete(c.cache, name)
}

func (c *TzCacheUnsafe) Contains(name string) bool {
	_, exists := c.cache[name]
	return exists
}

func (c *TzCacheUnsafe) Size() int {
	return len(c.cache)
}

func (c *TzCacheUnsafe) List() []string {
	names := make([]string, 0, len(c.cache))
	for name := range c.cache {
		names = append(names, name)
	}
	return names
}

func (c *TzCacheUnsafe) Clear() {
	c.cache = make(map[string]*time.Location)
}

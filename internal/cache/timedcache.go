package cache

import (
	"sync"
	"time"
)

// TimedCache caches data in a map[string]interface{}, is thread safe, and purges the cache after lifespan seconds.
type TimedCache struct {
	cache    map[string]interface{}
	lifespan float64
	makeTime time.Time
	mu       sync.Mutex
}

// NewTimedCache creates a new TimedCache with the given lifespan.
// A TimedCache should only ever be passed as a pointer value.
func NewTimedCache(lifespan float64) *TimedCache {
	return &TimedCache{
		cache:    make(map[string]interface{}),
		lifespan: lifespan,
		makeTime: time.Now(),
		mu:       sync.Mutex{},
	}
}

// checkTime checks the cache times and purges it if required.
// c.mu should be locked before calling this.
func (c *TimedCache) checkTime() {
	if time.Now().Sub(c.makeTime).Seconds() > c.lifespan {
		c.cache = make(map[string]interface{})
		c.makeTime = time.Now()
	}
}

// Get gets a value from the cache.
// If it doesn't hit the cache, the returned value will be an empty anonymous struct.
func (c *TimedCache) Get(key string) (value interface{}, hit bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.checkTime()
	if cached, ok := c.cache[key]; ok {
		return cached, true
	}
	return struct{}{}, false
}

// Set sets a value in the cache.
func (c *TimedCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.checkTime()
	c.cache[key] = value
}

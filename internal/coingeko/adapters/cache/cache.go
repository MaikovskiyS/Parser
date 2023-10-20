package cache

import (
	"errors"
	"testtask/internal/coingeko/model"

	"sync"
	"time"
)

var defaultExpiration = time.Second * 30
var cleanupInterval = time.Minute * 10

type cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	store             map[string]Item
}

func New() *cache {
	store := make(map[string]Item, 0)
	return &cache{
		store:             store,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}
}
func (c *cache) Set(key string, value model.Coin, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.store[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}
func (c *cache) Get(key string) (model.Coin, bool) {
	c.RLock()

	defer c.RUnlock()

	item, found := c.store[key]

	if !found {
		return model.Coin{}, false
	}

	if item.Expiration > 0 {

		// Если в момент запроса кеш устарел возвращаем nil
		if time.Now().UnixNano() > item.Expiration {
			return model.Coin{}, false
		}

	}

	return item.Value, true
}
func (c *cache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.store[key]; !found {
		return errors.New("key not found")
	}

	delete(c.store, key)

	return nil
}

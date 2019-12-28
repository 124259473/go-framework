package cache

import (
	"sync"
)

type LocalCache struct {
	Value map[string]string
}

var once sync.Once

var Cache *LocalCache

func GetInstance() *LocalCache {
	once.Do(func() {
		Cache = &LocalCache{}
	})
	return Cache
}

func (cache *LocalCache) Put(key string, value string) {
	if key == "" || value == "" {
		return
	}
	Cache.Value[key] = value
}

func (cache *LocalCache) Get(key string) string {
	if key == "" {
		return ""
	}
	return Cache.Value[key]
}

func (cache *LocalCache) Del(key string) {
	if key == "" {
		return
	}
	delete(Cache.Value, key)
}

func init() {
	value := make(map[string]string)
	Cache = GetInstance()
	Cache.Value = value
}

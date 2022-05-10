package cache

import (
	"log"
	"sync"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(ke string)
	DelOldest()
	Len() int
}

//默认允许占用的最大内存
type safeCache struct {
	m          sync.RWMutex
	cache      Cache
	nhit, nget int
}

func newSafeCache(cache Cache) *safeCache {
	return &safeCache{
		cache: cache,
	}
}

func (sc *safeCache) Set(key string, value interface{}) {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key, value)
}

func (sc *safeCache) Get(key string) interface{} {
	sc.m.RLock()
	defer sc.m.RUnlock()
	sc.nget++
	if sc.cache == nil {
		return nil
	}

	v := sc.cache.Get(key)
	if v != nil {
		log.Println("[ToueCache] hit")
		sc.nhit++
	}
	return v
}

func (sc *safeCache)stat() *Stat {
	sc.m.RLock()
	defer sc.m.RUnlock()
	return &Stat {
		NHit: sc.nhit,
		NGet: sc.nget,
	}
}

type Stat struct {
	NHit, NGet int
}


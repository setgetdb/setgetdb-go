package setgetdb

import "sync"

const MaxRecordCache = 10000000

type cacheRecord struct {
	Key string
	Value string
}

type CacheManager struct {
	recordManager *RecordManager
	cache map[uint32]cacheRecord
	cacheMutex *sync.Mutex
}

func NewCacheManager(recordManager *RecordManager) *CacheManager {
	return &CacheManager{recordManager, make(map[uint32]cacheRecord), &sync.Mutex{}}
}

func (c *CacheManager) Get(key string) (error, string)  {
	indexCache := c.hash(key)
	c.cacheMutex.Lock()
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		c.cacheMutex.Unlock()
		return nil, valueRetrievedByCache.Value
	}
	err, valueRetrievedByFile := c.recordManager.Get(key)
	c.cache[indexCache] = cacheRecord{key, valueRetrievedByFile }
	c.cacheMutex.Unlock()
	return err, valueRetrievedByFile
}

func (c *CacheManager) Set(key string, value string) error {
	indexCache := c.hash(key)
	c.cacheMutex.Lock()
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		c.cache[indexCache] = cacheRecord{key, value }
	}
	c.cacheMutex.Unlock()
	return c.recordManager.Set(key, value)
}

func (c *CacheManager) Delete(key string) error {
	indexCache := c.hash(key)
	c.cacheMutex.Lock()
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		delete(c.cache, indexCache)
	}
	c.cacheMutex.Unlock()
	return c.recordManager.Delete(key)
}

func (c *CacheManager) Close() error {
	return c.recordManager.Close()
}

func (c *CacheManager) hash(s string) uint32 {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (101 * hash) + int(s[i])
	}
	return uint32(hash) % MaxRecordCache
}
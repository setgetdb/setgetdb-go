package setgetdb

import "hash/fnv"

const MaxRecordCache = 10000000

type cacheRecord struct {
	Key string
	Value string
}

type CacheManager struct {
	recordManager *Recordmanager
	cache map[uint32]cacheRecord
}

func NewCacheManager(recordManager *Recordmanager) *CacheManager {
	return &CacheManager{recordManager, make(map[uint32]cacheRecord)}
}

func (c *CacheManager) Get(key string) (error, string)  {
	indexCache := c.hash(key)
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		return nil, valueRetrievedByCache.Value
	}
	err, valueRetrievedByFile := c.recordManager.Get(key)
	c.cache[indexCache] = cacheRecord{key, valueRetrievedByFile }
	return err, valueRetrievedByFile
}

func (c *CacheManager) Set(key string, value string) error {
	indexCache := c.hash(key)
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		c.cache[indexCache] = cacheRecord{key, value }
	}
	return c.recordManager.Set(key, value)
}

func (c *CacheManager) Delete(key string) error {
	indexCache := c.hash(key)
	valueRetrievedByCache := c.cache[indexCache]
	if valueRetrievedByCache.Value != "" && valueRetrievedByCache.Key == key {
		delete(c.cache, indexCache)
	}
	return c.recordManager.Delete(key)
}

func (c *CacheManager) Close() error {
	return c.recordManager.Close()
}

func (c *CacheManager) hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % MaxRecordCache
}
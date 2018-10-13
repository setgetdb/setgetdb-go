package setgetdb

const MaxRecordCache = 10000000

type cacheRecord struct {
	Key string
	Value string
}

type CacheManager struct {
	recordManager *RecordManager
	cache map[uint32]cacheRecord
}

func NewCacheManager(recordManager *RecordManager) *CacheManager {
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
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (101 * hash) + int(s[i])
	}
	return uint32(hash) % MaxRecordCache
}
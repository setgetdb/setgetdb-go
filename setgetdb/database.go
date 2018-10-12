package setgetdb

type Database struct {
	Name string
	cacheManager *CacheManager
}

func NewDatabase(prefixPath string, dbName string) *Database {
	path := prefixPath + dbName
	fileManager := NewFileManager(path)
	recordManager := NewRecordManager(fileManager)
	cacheManager := NewCacheManager(recordManager)
	return &Database{dbName, cacheManager}
}

func (d *Database) Get(key string) (error, string)  {
	return d.cacheManager.Get(key)
}

func (d *Database) Set(key string, value string) error {
	return d.cacheManager.Set(key, value)
}

func (d *Database) Delete(key string) error {
	return d.cacheManager.Delete(key)
}

func (d *Database) Close() error {
	return d.cacheManager.Close()
}

package setget

import (
	"hash/fnv"
)

const MaxNumberRecord = 100000000000
const SizeRecord = 1000
const PrefixPath = "./"

type Database struct {
	Name string
	fileManager *FileManager
}

func NewDatabase(dbName string) *Database {
	path := PrefixPath + dbName
	fileManager := NewFileManager(path)
	return &Database{dbName, fileManager}
}

func (d *Database) GetValue(key string) (error, string)  {
	from, to := d.getRangeBufferByKey(key)
	err, buffer := d.fileManager.Read(from, to)
	return err, string(buffer)
}

func (d *Database) SetByKey(key string, value string) error {
	from, _ := d.getRangeBufferByKey(key)
	buffer := make([]byte, SizeRecord)
	copy(buffer, []byte(value))
	return d.fileManager.Write(from, buffer)
}

func (d *Database) DeleteByKey(key string) error {
	from, to := d.getRangeBufferByKey(key)
	return d.fileManager.Delete(from, to)
}

func (d *Database) Close() error {
	return d.fileManager.Close()
}

func (d *Database) hash(s string) uint32 {
	const maxIndex = MaxNumberRecord / SizeRecord
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % maxIndex
}

func (d *Database) getRangeBufferByKey(key string) (uint32, uint32) {
	from := d.hash(key)
	to := from + SizeRecord
	return from, to
}

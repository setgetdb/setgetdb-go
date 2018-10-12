package setgetdb

import (
	"bytes"
	"hash/fnv"
)

const MaxNumberRecord = 100000000000
const SizeRecord = 1000

type Recordmanager struct {
	fileManager *FileManager
}

func NewRecordManager(fileManager *FileManager) *Recordmanager {
	return &Recordmanager{fileManager}
}

func (c *Recordmanager) Get(key string) (error, string)  {
	from, to := c.getRangeBufferByKey(key)
	err, buffer := c.fileManager.Read(from, to)
	return err, string(bytes.Trim(buffer, "\u0000"))
}

func (c *Recordmanager) Set(key string, value string) error {
	from, _ := c.getRangeBufferByKey(key)
	buffer := make([]byte, SizeRecord)
	copy(buffer, []byte(value))
	return c.fileManager.Write(from, buffer)
}

func (c *Recordmanager) Delete(key string) error {
	from, to := c.getRangeBufferByKey(key)
	return c.fileManager.Delete(from, to)
}

func (c *Recordmanager) Close() error {
	return c.fileManager.Close()
}

func (c *Recordmanager) hash(s string) uint32 {
	const maxIndex = MaxNumberRecord / SizeRecord
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % maxIndex
}

func (c *Recordmanager) getRangeBufferByKey(key string) (uint32, uint32) {
	from := c.hash(key)
	to := from + SizeRecord
	return from, to
}

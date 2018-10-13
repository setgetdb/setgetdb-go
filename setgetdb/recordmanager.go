package setgetdb

import (
	"bytes"
)

const SizeRecord = 1000

type RecordManager struct {
	fileManager *FileManager
}

func NewRecordManager(fileManager *FileManager) *RecordManager {
	return &RecordManager{fileManager}
}

func (c *RecordManager) Get(key string) (error, string)  {
	from, to := c.getRangeBufferByKey(key)
	err, buffer := c.fileManager.Read(from, to)
	return err, string(bytes.Trim(buffer, "\u0000"))
}

func (c *RecordManager) Set(key string, value string) error {
	from, _ := c.getRangeBufferByKey(key)
	buffer := make([]byte, SizeRecord)
	copy(buffer, []byte(value))
	return c.fileManager.Write(from, buffer)
}

func (c *RecordManager) Delete(key string) error {
	from, to := c.getRangeBufferByKey(key)
	return c.fileManager.Delete(from, to)
}

func (c *RecordManager) Close() error {
	return c.fileManager.Close()
}

func (c *RecordManager) hash(word string) uint32 {
	hash := 0
	lengthWord := len(word)
	for i := 0; i < lengthWord; i++ {
		hash = (101 * hash) + int(word[i])
	}
	return uint32(hash) * SizeRecord
}

func (c *RecordManager) getRangeBufferByKey(key string) (uint32, uint32) {
	from := c.hash(key)
	to := from + SizeRecord
	return from, to
}

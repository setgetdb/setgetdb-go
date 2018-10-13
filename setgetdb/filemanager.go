package setgetdb

import (
	"os"
	"sync"
)

type FileManager struct {
	path string
	fileMutex *sync.Mutex
	file *os.File
}

func NewFileManager(path string) *FileManager {
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	return &FileManager{path, &sync.Mutex{},file}
}

func (f *FileManager) Write(from uint32, buffer []byte) error {
	f.fileMutex.Lock()
	f.file.Seek(int64(from), 0)
	_, err := f.file.Write(buffer)
	f.file.Sync()
	f.fileMutex.Unlock()
	return err
}

func (f *FileManager) Read(from uint32, to uint32) (error, []byte) {
	f.fileMutex.Lock()
	result := make([]byte, to-from)
	f.file.Seek(int64(from), 0)
	_, err := f.file.Read(result)
	f.fileMutex.Unlock()
	return err, result
}

func (f *FileManager) Delete(from uint32, to uint32) error {
	f.fileMutex.Lock()
	f.file.Seek(int64(from), 0)
	_, err := f.file.Write(make([]byte, to-from))
	f.file.Sync()
	f.fileMutex.Unlock()
	return err
}

func (f *FileManager) Close() error {
	f.fileMutex.Lock()
	err := f.file.Close()
	f.fileMutex.Unlock()
	return err
}


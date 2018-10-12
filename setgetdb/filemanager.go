package setgetdb

import (
	"os"
)

type FileManager struct {
	path string
	file *os.File
}

func NewFileManager(path string) *FileManager {
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	return &FileManager{path, file}
}

func (f *FileManager) Write(from uint32, buffer []byte) error {
	f.file.Seek(int64(from), 0)
	_, err := f.file.Write(buffer)
	f.file.Sync()
	return err
}

func (f *FileManager) Read(from uint32, to uint32) (error, []byte) {
	result := make([]byte, to-from)
	f.file.Seek(int64(from), 0)
	_, err := f.file.Read(result)
	return err, result
}

func (f *FileManager) Delete(from uint32, to uint32) error {
	f.file.Seek(int64(from), 0)
	_, err := f.file.Write(make([]byte, to-from))
	f.file.Sync()
	return err
}

func (f *FileManager) Close() error {
	return f.file.Close()
}


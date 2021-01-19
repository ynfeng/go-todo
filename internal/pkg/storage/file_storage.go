package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileStorage struct {
	dir      string
	fileName string
}

func (storage FileStorage) Append(item interface{}) error {
	file, err := openFile(storage)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer, string(bytes))
	if err != nil {
		return err
	}
	return writer.Flush()
}

func (storage FileStorage) All(factory ItemFactory) ([]interface{}, error) {
	file, err := openFile(storage)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]interface{}, 0)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		item := factory()
		err = json.Unmarshal([]byte(line), &item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func (storage FileStorage) replaceAll(items ...interface{}) error {
	file, err := truncFile(storage)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(writer, string(bytes))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func openFile(storage FileStorage) (*os.File, error) {
	ensureDirExists(storage.dir)
	filePath := storage.dir + storage.fileName
	return os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
}

func truncFile(storage FileStorage) (*os.File, error) {
	ensureDirExists(storage.dir)
	filePath := storage.dir + storage.fileName
	return os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
}

func ensureDirExists(dir string) {
	_ = os.Mkdir(dir, 0777)
}

func NewFileStorage(dir string, fileName string) *FileStorage {
	return &FileStorage{dir: dir, fileName: fileName}
}

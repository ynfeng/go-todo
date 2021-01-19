package todolist

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type FileStorage struct {
	dir      string
	fileName string
}

func (storage FileStorage) Append(item Item) {
	file, err := openFile(storage)
	if err != nil {
		log.Fatal("can't open data file.")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	bytes, err := json.Marshal(item)
	if err != nil {
		log.Fatal("json serialization error")
	}
	_, err = fmt.Fprintln(writer, string(bytes))
	if err != nil {
		log.Fatal("write data error")
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal("flush data error")
	}
}

func (storage FileStorage) All() []Item {
	file, err := openFile(storage)
	if err != nil {
		log.Fatal("can't open data file.")
	}
	defer file.Close()

	result := make([]Item, 0)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		item := Item{}
		err = json.Unmarshal([]byte(line), &item)
		if err != nil {
			log.Fatal("json deserialization error")
		}
		result = append(result, item)
	}
	return result
}

func (storage FileStorage) replaceAll(items ...Item) {
	file, err := truncFile(storage)
	if err != nil {
		log.Fatal("can't open data file.")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			log.Fatal("json serialization error")
		}
		_, err = fmt.Fprintln(writer, string(bytes))
		if err != nil {
			log.Fatal("write data error")
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal("flush data error")
	}
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

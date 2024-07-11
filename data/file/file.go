package file

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Pair struct {
	Key   string     `json:"key"`
	Value *ValueInfo `json:"value"`
}

type ValueInfo struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type FileDb struct {
	filepath string
	mode     int64
}

const (
	FilePathLocal  = "/home/sunlie/Project/SunCache/data/file/db_pair.json"
	FilePathStream = "/home/sunlie/Project/SunCache/data/file/db_pairArray.json"
)

const (
	ModeLocal  = 1 << iota
	ModeStream = 1 << iota
)

func NewFileDb(mode int64) (fileDb *FileDb) {
	fileDb = &FileDb{
		mode: mode,
	}
	return
}

func (f *FileDb) Get(key string) (value []byte, err error) {
	var res *ValueInfo
	if f.mode == ModeLocal {
		f.filepath = FilePathLocal
		var file *os.File
		file, err = os.Open(f.filepath)
		if err != nil {
			return
		}
		defer file.Close()

		var byteValue []byte
		byteValue, err = io.ReadAll(file)
		if err != nil {
			log.Fatalf("Read file %s error: %v", f.filepath, err)
			return
		}

		var data map[string]*ValueInfo
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			log.Fatalf("Unmarshal json error: %v", err)
		}
		res = data[key]
	} else if f.mode == ModeStream {
		f.filepath = FilePathStream
		var file *os.File
		file, err = os.Open(f.filepath)
		if err != nil {
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)

		decoder.Token()
		for decoder.More() {
			var pair Pair
			err = decoder.Decode(&pair)
			if err != nil {
				if err == io.EOF {
					return nil, nil
				}
				return
			}
			if pair.Key == key {
				res = pair.Value
				break
			}
		}
	}

	value, err = json.Marshal(res)

	return
}

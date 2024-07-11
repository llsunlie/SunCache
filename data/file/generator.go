package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/google/uuid"
)

func generatePair() (pair *Pair) {
	uuid := uuid.New()
	key := uuid.String()

	valueInfo := &ValueInfo{
		Uuid: key,
		Name: namesgenerator.GetRandomName(0),
		Age:  rand.Intn(100),
	}
	pair = &Pair{
		Key:   key,
		Value: valueInfo,
	}
	return
}

func GeneratePair(count int) {
	filepath := FilePathLocal
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	pairs := make(map[string]*ValueInfo)
	for i := 0; i < count; i++ {
		pair := generatePair()
		pairs[pair.Key] = pair.Value
	}

	jsonByte, err := json.Marshal(pairs)
	if err != nil {
		log.Fatalf("Failed to marshal json: %v", err)
	}
	jsonString := string(jsonByte)
	file.WriteString(jsonString)
}

func GeneratePairArray(count int) {
	filepath := FilePathStream
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	pairs := make([]*Pair, 0)
	for i := 0; i < count; i++ {
		pair := generatePair()
		pairs = append(pairs, pair)
	}

	jsonByte, err := json.Marshal(pairs)
	if err != nil {
		log.Fatalf("Failed to marshal json: %v", err)
	}
	jsonString := string(jsonByte)
	file.WriteString(jsonString)
}

func ReadPair() {
	filepath := FilePathLocal
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Open file %s error: %v", filepath, err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read file %s error: %v", filepath, err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Unmarshal json error: %v", err)
	}
	log.Printf("data: %+v", data)
}

func ReadPairArray() {
	filepath := FilePathStream
	var file *os.File
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	decoder.Token()
	for decoder.More() {
		var pair Pair
		err := decoder.Decode(&pair)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", pair.Key, pair.Value)
	}
	decoder.Token()

}

func ReadPairKeys() (keys []string) {
	filepath := FilePathLocal
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Open file %s error: %v", filepath, err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read file %s error: %v", filepath, err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Unmarshal json error: %v", err)
	}
	keys = make([]string, 0)
	for key := range data {
		keys = append(keys, key)
	}
	return
}

func ReadPairArrayKeys() (keys []string) {
	filepath := FilePathStream
	var file *os.File
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	keys = make([]string, 0)

	decoder := json.NewDecoder(file)

	decoder.Token()
	for decoder.More() {
		var pair Pair
		err := decoder.Decode(&pair)
		if err != nil {
			log.Fatal(err)
		}
		keys = append(keys, pair.Key)
	}
	return
}

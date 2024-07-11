package file_test

import (
	"SunCache/data/file"
	"fmt"
	"testing"
)

func TestLocalGet(t *testing.T) {
	db := file.NewFileDb(file.ModeLocal)
	value, err := db.Get("0015a715-f784-44d9-b381-f064627c40c3")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("value: %v\n", value)
}

func TestStreamGet(t *testing.T) {
	db := file.NewFileDb(file.ModeStream)
	value, err := db.Get("43f18d44-758a-4213-a07b-e9bbe8f02f70")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("value: %v\n", value)
}

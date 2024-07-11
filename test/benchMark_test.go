package test_test

import (
	"SunCache/data/file"
	"math/rand"
	"net/http"
	"testing"
)

func BenchmarkCache_FileModeLocal(b *testing.B) {
	keys := file.ReadPairKeys()
	count := len(keys)
	for i := 0; i < b.N; i++ {
		key := keys[rand.Intn(count)]
		_, err := http.Get("http://localhost:8300/api?member=userInfo&key=" + key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCache_FileModeStream(b *testing.B) {
	keys := file.ReadPairArrayKeys()
	count := len(keys)
	for i := 0; i < b.N; i++ {
		key := keys[rand.Intn(count)]
		_, err := http.Get("http://localhost:8300/api?member=userInfo&key=" + key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

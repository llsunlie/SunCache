package test_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestCache(t *testing.T) {
	res, err := http.Get("http://localhost:8300/api?member=userInfo&key=43f18d44-758a-4213-a07b-e9bbe8f02f70")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("content: %v\n", string(content))
}

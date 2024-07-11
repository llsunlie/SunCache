package file_test

import (
	"SunCache/data/file"
	"testing"
)

func TestGeneratePair(t *testing.T) {
	file.GeneratePair(10000)
}

func TestReadPair(t *testing.T) {
	file.ReadPair()
}

func TestGeneratePairArray(t *testing.T) {
	file.GeneratePairArray(10000)
}

func TestReadPairArray(t *testing.T) {
	file.ReadPairArray()
}

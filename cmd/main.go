package main

import (
	"github.com/naturali/mapit/merger"
)

func init() {
	merger.SamplingSize = 1000000
}

func main() {
	kvsArray := [][]merger.KV{}
	arraySize := 8
	for i := 0; i < 8; i++ {
		kvsArray = append(kvsArray, merger.MakeSortedRandKVArray(arraySize))
		arraySize *= 4
	}

}

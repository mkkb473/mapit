package main

import (
	"fmt"

	"github.com/naturali/mapit/merger"
)

func init() {
	merger.SamplingSize = 4000000
}

func main() {
	kvsArray := [][]merger.KV{}
	arraySize := 8
	for i := 0; i < 8; i++ {
		kvsArray = append(kvsArray, merger.MakeSortedRandKVArray(arraySize))
		arraySize *= 4
	}

	retArrayP := &kvsArray[0]
	for i := 1; i < len(kvsArray); i++ {
		retArrayP = merger.MergeTwoArries(retArrayP, &kvsArray[i])
	}

	for _, iter := range *retArrayP {
		fmt.Println(iter)
	}
	fmt.Println(len(*retArrayP))
}

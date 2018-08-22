package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/naturali/mapit/merger"
	"github.com/qiniu/log"
)

func init() {
	fs := flag.NewFlagSet("name", flag.ExitOnError)
	samplingBound := fs.Int("sampling-bound", 1000000, "sampling size")
	bytesSize := fs.Int("bytes-size", 4, "byte array size")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if *samplingBound < 131072 {
		fmt.Println("bytes size should be greater than 131072 (8*(4**7))")
		os.Exit(1)
	}
	if *samplingBound > 200000000 {
		fmt.Println("You don't want to run out of memory")
		os.Exit(1)
	}
	if *bytesSize < 4 {
		fmt.Println("Beta version, only size greater than 4 is supported now;")
		os.Exit(1)
	}
	merger.SamplingBound = *samplingBound
	merger.BytesSize = *bytesSize
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
		for num := iter.Value.Next(); num != nil; num = iter.Value.Next() {
			fmt.Print(num, " ")
		}
	}
	fmt.Println(len(*retArrayP))

}

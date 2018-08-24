package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/naturali/mapit/merger"
	"github.com/qiniu/log"
)

func init() {
	fs := flag.NewFlagSet("name", flag.ExitOnError)
	samplingBound := fs.Int("sampling-bound", 1000000, "sampling size")
	bytesSize := fs.Int("bytes-size", 4, "byte array size")
	iteratorSize := fs.Int("iter-size", 4, "iterator data size")
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
		fmt.Println("Beta version, only size greater than 3 is supported now;")
		os.Exit(1)
	}
	if *iteratorSize > 100 || *iteratorSize < 1 {
		fmt.Println("Iterator size should be greater than 0 and less than 100")
		os.Exit(1)
	}
	merger.SamplingBound = *samplingBound
	merger.BytesSize = *bytesSize
}

func main() {
	kvsArray := [][]merger.KV{}
	arraySize := 8
	time1 := time.Now()
	fmt.Println("Generating random kvs array...")

	for i := 0; i < 8; i++ {
		kvsArray = append(kvsArray, merger.MakeSortedRandKVArray(arraySize))
		arraySize *= 4
	}
	/*
		for i := 0; i < 8; i++ {
			kvsArray = append(kvsArray, merger.MakeSortedRandKVArray(150000))
			arraySize *= 4
		}
	*/
	fmt.Println("Completed generation.")
	time2 := time.Now()

	fmt.Println("Start merging...")
	retArrayP := &kvsArray[0]
	for i := 1; i < len(kvsArray); i++ {
		retArrayP = merger.MergeTwoArries(retArrayP, &kvsArray[i])
	}
	fmt.Println("Completed merging.")
	time3 := time.Now()
	fmt.Println("Printing iterations....")
	for _, iter := range *retArrayP {
		for num := iter.Value.Next(); num != nil; num = iter.Value.Next() {
			fmt.Print(num, " ")
		}
	}
	fmt.Println(len(*retArrayP))
	time4 := time.Now()
	fmt.Printf("%v\n%v\n%v\n%v\n", time1, time2, time3, time4)
	fmt.Printf("Generating input:\t\n%vMerging arries:\t%v\nIteration:\t%v", time2.Sub(time1), time3.Sub(time2), time4.Sub(time3))
}

package merger

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sort"
	"time"
)

var (
	// SamplingBound is the upper bound of int->[]byte for keys
	// should be greater than 8 * (4**7)
	SamplingBound int

	// BytesSize is the size of bytes array as Key
	BytesSize int
)

// KV is Key-Value store data structure
type KV struct {
	Key   []byte
	Value Iterator
}

// Iterator pretty weak & slow iterator
type Iterator struct {
	Data []int
	I    int
}

// Next gives next element in iterator
func (iter *Iterator) Next() interface{} {
	if iter.I == len(iter.Data) {
		return nil
	}
	ret := iter.Data[iter.I]
	iter.I++
	return ret
}

// MakeSortedRandKVArray returns an array of Key-Value pair with random sampled keys
// KV's Key is parsed from randomly generated int32 with BigEndian
func MakeSortedRandKVArray(keyNum int) []KV {
	kvArr := []KV{}
	rand.Seed(time.Now().UnixNano())
	keys := rand.Perm(SamplingBound)[:keyNum]
	sort.Ints(keys)
	for _, value := range keys {
		bs := make([]byte, BytesSize)
		binary.BigEndian.PutUint32(bs, uint32(value))
		rand.Seed(time.Now().UnixNano())
		data := rand.Perm(20)[:rand.Intn(4)+1]
		sort.Ints(data)
		kvArr = append(kvArr, KV{
			Key: bs,
			Value: Iterator{
				Data: data,
				I:    0,
			},
		})
	}
	return kvArr
}

// MergeTwoArries merges two sorted KV arries
func MergeTwoArries(A, B *[]KV) *[]KV {
	var C []KV
	// A should always be shorter than B
	if len(*A) > len(*B) {
		A, B = B, A
	}

	// Same as merge two sorted list, O(n+4n) runtime
	var headA, headB int
	for headA != len(*A) && headB != len(*B) {
		bytesCompare := bytes.Compare((*A)[headA].Key, (*B)[headB].Key)
		if bytesCompare < 0 {
			C = append(C, (*A)[headA])
			headA++
		} else if bytesCompare == 0 {
			C = append(C, (*A)[headA])
			headA++
			headB++
		} else {
			C = append(C, (*B)[headB])
			headB++
		}
	}
	if headA != len(*A) {
		for headA < len(*A) {
			C = append(C, (*A)[headA])
			headA++
		}
	}
	if headB != len(*B) {
		for headB < len(*B) {
			C = append(C, (*B)[headB])
			headB++
		}
	}
	return &C
}

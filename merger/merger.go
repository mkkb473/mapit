package merger

import (
	"encoding/binary"
	"math/rand"
	"sort"
	"time"
)

var (
	// ArrayValueLength is the length of Ierator.Data as value in KV
	ArrayValueLength int
	// SamplingSize is the upper bound of int->[]byte for keys
	SamplingSize int

	fakeIterData []int
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
	iter.I++
	return iter.Data[iter.I]
}

func MakeSortedRandKVArray(keyNum int) []KV {
	kvArr := []KV{}
	rand.Seed(time.Now().UnixNano())
	keys := rand.Perm(SamplingSize)[:keyNum]
	sort.Ints(keys)
	for _, value := range keys {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(value))
		rand.Seed(time.Now().UnixNano())
		kvArr = append(kvArr, KV{
			Key: bs,
			Value: Iterator{
				Data: rand.Perm(20)[:rand.Intn(4)+1],
				I:    0,
			},
		})
	}
	return kvArr
}

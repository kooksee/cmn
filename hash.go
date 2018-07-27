package cmn

import (
	"hash"
	"io"
	"encoding/binary"
)

var Hash = myHash{}

type myHash struct{}

type jumpHash struct{}

func (myHash) JumpHash() jumpHash {
	return jumpHash{}
}

func (jumpHash) Hash(key uint64, buckets int) int {
	var b, j int64
	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int(b)
}

func (m jumpHash) HashInt(key uint64, buckets int, length int) []int {
	res := make([]int, 0, length)

	if buckets <= 0 {
		buckets = 1
	}

	for i := 0; i < length; {
		h := m.Hash(key, buckets)
		for j := 0; j < len(res); j++ {
			if res[j] == h {
				h = -1
				break
			}
			if res[j] > h {
				res[j], h = h, res[j]
			}
		}
		if h >= 0 {
			res = append(res, h)
			i++
		}
		key++

	}
	return res
}

func (m jumpHash) HashString(key string, buckets int, length int, keyHasher hash.Hash64) []int {
	keyHasher.Reset()

	_, err := io.WriteString(keyHasher, key)
	if err != nil {
		panic(err)
	}
	return m.HashInt(keyHasher.Sum64(), buckets, length)
}

type murmurHash struct{}

func (myHash) MurmurHash() murmurHash {
	return murmurHash{}
}

// Hash return hash of the given data.
func (murmurHash) Hash(data []byte, seed uint32) uint32 {
	// Similar to murmur hash
	const (
		m = uint32(0xc6a4a793)
		r = uint32(24)
	)
	var (
		h = seed ^ (uint32(len(data)) * m)
		i int
	)

	for n := len(data) - len(data)%4; i < n; i += 4 {
		h += binary.LittleEndian.Uint32(data[i:])
		h *= m
		h ^= (h >> 16)
	}

	switch len(data) - i {
	default:
		panic("not reached")
	case 3:
		h += uint32(data[i+2]) << 16
		fallthrough
	case 2:
		h += uint32(data[i+1]) << 8
		fallthrough
	case 1:
		h += uint32(data[i])
		h *= m
		h ^= (h >> r)
	case 0:
	}

	return h
}

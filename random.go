package cmn

import (
	crand "crypto/rand"
	mrand "math/rand"
	"sync"
	"time"
)

var Rand = myRand{
	strChars: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", // 62 characters
}

type myRand struct {
	strChars string
}

// pseudo random number generator.
// seeded with OS randomness (crand)
var prng struct {
	sync.Mutex
	*mrand.Rand
}

func reset() {
	b := cRandBytes(8)
	var seed uint64
	for i := 0; i < 8; i++ {
		seed |= uint64(b[i])
		seed <<= 8
	}
	prng.Lock()
	prng.Rand = mrand.New(mrand.NewSource(int64(seed)))
	prng.Unlock()
}

func init() {
	reset()
}

// Constructs an alphanumeric string of given length.
// It is not safe for cryptographic usage.
func (m myRand) RandStr(length int) string {
	chars := []byte{}
MAIN_LOOP:
	for {
		val := m.RandInt63()
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) // rightmost 6 bits
			if v >= 62 { // only 62 characters in strChars
				val >>= 6
				continue
			} else {
				chars = append(chars, m.strChars[v])
				if len(chars) == length {
					break MAIN_LOOP
				}
				val >>= 6
			}
		}
	}

	return string(chars)
}

// It is not safe for cryptographic usage.
func (m myRand) RandUint16() uint16 {
	return uint16(m.RandUint32() & (1<<16 - 1))
}

// It is not safe for cryptographic usage.
func (m myRand) RandUint32() uint32 {
	prng.Lock()
	u32 := prng.Uint32()
	prng.Unlock()
	return u32
}

// It is not safe for cryptographic usage.
func (m myRand) RandUint64() uint64 {
	return uint64(m.RandUint32())<<32 + uint64(m.RandUint32())
}

// It is not safe for cryptographic usage.
func (m myRand) RandUint() uint {
	prng.Lock()
	i := prng.Int()
	prng.Unlock()
	return uint(i)
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt16() int16 {
	return int16(m.RandUint32() & (1<<16 - 1))
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt32() int32 {
	return int32(m.RandUint32())
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt64() int64 {
	return int64(m.RandUint64())
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt() int {
	prng.Lock()
	i := prng.Int()
	prng.Unlock()
	return i
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt31() int32 {
	prng.Lock()
	i31 := prng.Int31()
	prng.Unlock()
	return i31
}

// It is not safe for cryptographic usage.
func (m myRand) RandInt63() int64 {
	prng.Lock()
	i63 := prng.Int63()
	prng.Unlock()
	return i63
}

// Distributed pseudo-exponentially to test for various cases
// It is not safe for cryptographic usage.
func (m myRand) RandUint16Exp() uint16 {
	bits := m.RandUint32() % 16
	if bits == 0 {
		return 0
	}
	n := uint16(1 << (bits - 1))
	n += uint16(m.RandInt31()) & ((1 << (bits - 1)) - 1)
	return n
}

// Distributed pseudo-exponentially to test for various cases
// It is not safe for cryptographic usage.
func (m myRand) RandUint32Exp() uint32 {
	bits := m.RandUint32() % 32
	if bits == 0 {
		return 0
	}
	n := uint32(1 << (bits - 1))
	n += uint32(m.RandInt31()) & ((1 << (bits - 1)) - 1)
	return n
}

// Distributed pseudo-exponentially to test for various cases
// It is not safe for cryptographic usage.
func (m myRand) RandUint64Exp() uint64 {
	bits := m.RandUint32() % 64
	if bits == 0 {
		return 0
	}
	n := uint64(1 << (bits - 1))
	n += uint64(m.RandInt63()) & ((1 << (bits - 1)) - 1)
	return n
}

// It is not safe for cryptographic usage.
func (m myRand) RandFloat32() float32 {
	prng.Lock()
	f32 := prng.Float32()
	prng.Unlock()
	return f32
}

// It is not safe for cryptographic usage.
func (m myRand) RandTime() time.Time {
	return time.Unix(int64(m.RandUint64Exp()), 0)
}

// RandBytes returns n random bytes from the OS's source of entropy ie. via crypto/rand.
// It is not safe for cryptographic usage.
func (m myRand) RandBytes(n int) []byte {
	// cRandBytes isn't guaranteed to be fast so instead
	// use random bytes generated from the internal PRNG
	bs := make([]byte, n)
	for i := 0; i < len(bs); i++ {
		bs[i] = byte(m.RandInt() & 0xFF)
	}
	return bs
}

// RandIntn returns, as an int, a non-negative pseudo-random number in [0, n).
// It panics if n <= 0.
// It is not safe for cryptographic usage.
func (m myRand) RandIntn(n int) int {
	prng.Lock()
	i := prng.Intn(n)
	prng.Unlock()
	return i
}

// RandPerm returns a pseudo-random permutation of n integers in [0, n).
// It is not safe for cryptographic usage.
func (m myRand) RandPerm(n int) []int {
	prng.Lock()
	perm := prng.Perm(n)
	prng.Unlock()
	return perm
}

func (m myRand) Rand32(max uint32) uint32 {
	if max == 0 {
		return 0
	}
	mrand.Seed(time.Now().Unix())
	return mrand.Uint32() % max
}

// 生成count个[start,end)结束的不重复的随机数
func (m myRand) GenRandom(start int, end int, count int) map[int]bool {
	nums := make(map[int]bool)

	// 范围检查
	if end < start || (end-start) < count {
		return nums
	}

	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {

		// 生成随机数
		num := r.Intn(end-start) + start
		if nums[num] {
			continue
		}
		nums[num] = true
	}

	return nums
}

// NOTE: This relies on the os's random number generator.
// For real security, we should salt that with some seed.
// See github.com/tendermint/go-crypto for a more secure reader.
func cRandBytes(numBytes int) []byte {
	b := make([]byte, numBytes)
	_, err := crand.Read(b)
	if err != nil {
		Err.MustNotErr(err)
	}
	return b
}

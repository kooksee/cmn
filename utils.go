package cmn

import (
	"math"
	"encoding/binary"
	"fmt"
	"os"
	"bytes"
	"reflect"
	"unsafe"
	"strconv"
)

func F(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func BytesTrimSpace(bs []byte) []byte {
	for i, b := range bs {
		if b != 0 {
			bs = bs[i:]
			break
		}
	}

	for i, b := range bs {
		if b == 0 {
			bs = bs[:i]
			break
		}
	}

	return bs
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

// 使用二进制存储整形
func Int64ToByte(x int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(x))
	return b
}

func ByteToInt64(x []byte) int64 {
	return int64(binary.BigEndian.Uint64(x))
}

func StructSortMarshal(s interface{}) ([]byte, error) {
	s1, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	b := map[string]interface{}{}
	if err := json.Unmarshal(s1, &b); err != nil {
		return nil, err
	}
	b1, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return b1, nil
}

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func If(b bool, trueVal, falseVal interface{}) interface{} {
	if b {
		return trueVal
	}
	return falseVal
}

// 范围判断 min <= v <= max
func Between(v, min, max []byte) bool {
	return bytes.Compare(v, min) >= 0 && bytes.Compare(v, max) <= 0
}

// 复制数组
func CopyBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

// 使用二进制存储整形
func IntToByte(x int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(x))
	return b
}

func ByteToInt(x []byte) int {
	return int(binary.BigEndian.Uint64(x))
}

// S2b converts string to a byte slice without memory allocation.
// "abc" -> []byte("abc")
func S2b(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// B2s converts byte slice to a string without memory allocation.
// []byte("abc") -> "abc" s
func B2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// B2ds return a Digit string of v
// v (8-byte big endian) -> uint64(123456) -> "123456".
func B2ds(v []byte) string {
	return strconv.FormatUint(binary.BigEndian.Uint64(v), 10)
}

// Btoi return an int64 of v
// v (8-byte big endian) -> uint64(123456).
func B2i(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}

// DS2i returns uint64 of Digit string
// v ("123456") -> uint64(123456).
func DS2i(v string) uint64 {
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return uint64(0)
	}
	return i
}

// Itob returns an 8-byte big endian representation of v
// v uint64(123456) -> 8-byte big endian.
func I2b(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// DS2b returns an 8-byte big endian representation of Digit string
// v ("123456") -> uint64(123456) -> 8-byte big endian.
func DS2b(v string) []byte {
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return []byte("")
	}
	return I2b(i)
}

// BConcat concat a list of byte
func BConcat(slices ... []byte) []byte {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	tmp := make([]byte, totalLen)
	var i int
	for _, s := range slices {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func BMap(m [][]byte, fn func(i int, k []byte) []byte) [][]byte {
	for i, d := range m {
		m[i] = fn(i, d)
	}
	return m
}

func P(a ...interface{}) {
	fmt.Println(a...)
}

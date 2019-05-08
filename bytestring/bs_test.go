package bytestring

import (
	"fmt"
	"math/rand"
	"testing"
)

var byte16 [][]byte
var byte256 [][]byte
var byte1024 [][]byte

var string16 []string
var string256 []string
var string1024 []string

func init()  {
	for i := 0; i < 100; i++ {
		byte16 = append(byte16, genByte(16))
		byte256 = append(byte256, genByte(256))
		byte1024 = append(byte1024, genByte(1024))

		string16 = append(string16, genString(16))
		string256 = append(string256, genString(256))
		string1024 = append(string1024, genString(1024))
	}
}

func genByte(len int) []byte {
	buf := make([]byte, len, len)
	for i := 0; i < len; i++ {
		buf[i] = 'A' + byte(rand.Intn(62))
	}
	return buf
}

func genString(len int) string {
	buf := make([]byte, len, len)
	for i := 0; i < len; i++ {
		buf[i] = 'A' + byte(rand.Intn(62))
	}
	return string(buf)
}

func TestBytesToString(t *testing.T) {
	s := "hello，我的中国心"
	bs := StringToBytes(s)
	s1 := BytesToString(bs)
	bs1 := StringToBytes(s1)

	fmt.Println(bs)
	fmt.Println(s1)
	fmt.Printf("%p, %p", bs, bs1)
}

func Benchmark_BytesToString16(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte16 {
			s = BytesToString(buf)
		}
	}
	s = s
}

func Benchmark_String16(b *testing.B)  {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte16 {
			s = string(buf)
		}
	}
	s = s
}

func Benchmark_BytesToString256(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte256 {
			s = BytesToString(buf)
		}
	}
	s = s
}

func Benchmark_String256(b *testing.B)  {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte256 {
			s = string(buf)
		}
	}
	s = s
}

func Benchmark_BytesToString1024(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte1024 {
			s = BytesToString(buf)
		}
	}
	s = s
}

func Benchmark_String1024(b *testing.B)  {
	var s string
	for i := 0; i < b.N; i++ {
		for _, buf := range byte1024 {
			s = string(buf)
		}
	}
	s = s
}


func Benchmark_StringToBytes16(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string16 {
			bs = StringToBytes(s)
		}
	}
	bs = bs
}

func Benchmark_Byte16(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string16 {
			bs = []byte(s)
		}
	}
	bs = bs
}

func Benchmark_StringToBytes256(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string256 {
			bs = StringToBytes(s)
		}
	}
	bs = bs
}

func Benchmark_Byte256(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string256 {
			bs = []byte(s)
		}
	}
	bs = bs
}

func Benchmark_StringToBytes1024(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string1024 {
			bs = StringToBytes(s)
		}
	}
	bs = bs
}

func Benchmark_Byte1024(b *testing.B) {
	var bs []byte
	for i := 0; i < b.N; i++ {
		for _, s := range string1024 {
			bs = []byte(s)
		}
	}
	bs = bs
}
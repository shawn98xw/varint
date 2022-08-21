package main

const maxVarintBytes = 10 // maximum length of a varint

// EncodeVarint returns the varint encoding of x.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
// Not used by the package itself, but helpful to clients
// wishing to use the same encoding.
// 魔数
// 0x80 => 0000000010000000
// 0x7f => 0000000001111111
func EncodeVarint(x uint64) []byte {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		// 首位记 1, 写入原始数字从低位始的 7 个 bit
		buf[n] = 0x80 | uint8(x&0x7F)
		// 移走记录过的 7 位
		x >>= 7
	}
	// 剩余不足 7 位的部分直接以 8 位形式存下来，故首位为 0
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

// DecodeVarint reads a varint-encoded integer from the slice.
// It returns the integer and the number of bytes consumed, or
// zero if there is not enough.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
func DecodeVarint(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		// 弃首位取 7 位并加回 x
		x |= (b & 0x7F) << shift
		// 首位为 0
		if (b & 0x80) == 0 {
			return x, n
		}
	}

	// The number is too large to represent in a 64-bit value.
	return 0, 0
}

func main() {
	println("hello")
	res := EncodeVarint(123)
	a, b := DecodeVarint(res)
	// println(res)
	println(a)
	println(b)
}

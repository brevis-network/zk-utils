package circuits

func Pad101(data []byte) []uint64 {
	return Bytes2Uint64s(Pad101Bytes(data))
}

func Pad101Bytes(data []byte) []byte {
	miss := 136 - len(data)%136
	if len(data)%136 == 0 {
		miss = 136
	}
	data = append(data, 1)
	for i := 0; i < miss-1; i++ {
		data = append(data, 0)
	}
	data[len(data)-1] ^= 0x80
	return data
}

func Bytes2BlockBits(bytes []byte) (bits []uint8) {
	if len(bytes)%136 != 0 {
		panic("invalid length")
	}
	return Bytes2Bits(bytes)
}

func Bytes2Uint64(bytes []byte) uint64 {
	if len(bytes) != 8 {
		panic("bytes len must be 8")
	}
	ret := uint64(0)
	for i := 0; i < 8; i++ {
		ret |= uint64(bytes[i]) << (i * 8)
	}
	return ret
}

func Bytes2Bits(bytes []byte) (bits []uint8) {
	if len(bytes)%8 != 0 {
		panic("invalid length")
	}
	for i := 0; i < len(bytes); i++ {
		bits = append(bits, byte2Bits(bytes[i])...)
	}
	return
}

// bytes2Bits outputs bits in little-endian
func byte2Bits(b byte) (bits []uint8) {
	for i := 0; i < 8; i++ {
		bits = append(bits, (uint8(b)>>i)&1)
	}
	return
}

func Bytes2Uint64s(bytes []byte) []uint64 {
	if len(bytes)%8 != 0 {
		panic("bytes len must be multiple of 8")
	}
	ret := []uint64{}
	for i := 0; i < len(bytes); i += 8 {
		ret = append(ret, Bytes2Uint64(bytes[i:i+8]))
	}
	return ret
}

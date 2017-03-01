package murmurhash

func rotl32(x uint32, r uint8) uint32 {
	return x<<r | x>>(32-r)
}

func fmix32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func MurmurHash3_x86_32(key []byte, seed uint32) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	var k uint32
	h, l := seed, uint32(len(key))

	for len(key) >= 4 {
		k = uint32(key[0]) |
			uint32(key[1])<<8 |
			uint32(key[2])<<16 |
			uint32(key[3])<<24
		k *= c1
		k = rotl32(k, r1)
		k *= c2

		h ^= k
		h = rotl32(h, r2)
		h = h*m + n

		key = key[4:]
	}

	k = 0
	switch len(key) & 3 {
	case 3:
		k ^= uint32(key[2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[0])
		k *= c1
		k = rotl32(k, r1)
		k *= c2
		h ^= k
	}

	h ^= l
	h = fmix32(h)
	return h
}

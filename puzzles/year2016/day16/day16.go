package day16

var (
	diskSize = 272
	one      = byte('1')
	zero     = byte('0')
)

func reverseInvert(b []byte) {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
	for i, v := range b {
		b[i] = inverse(v)
	}
}

func inverse(b byte) byte {
	if b == zero {
		return one
	}
	return zero
}

func generate(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	a = append(a, zero)
	reverseInvert(b)
	return append(a, b...)
}

func checksum(b []byte) []byte {
	var idx int
	for i := 0; i < len(b)-1; i += 2 {
		if b[i] == b[i+1] {
			b[idx] = one
		} else {
			b[idx] = zero
		}
		idx++
	}
	b = b[:idx]
	if len(b)%2 == 0 {
		return checksum(b)
	}
	return b
}

func fillDisk(start []byte, size int) []byte {
	data := make([]byte, 0, size)
	data = append(data, start...)
	for len(data) < size {
		data = generate(data)
	}
	return data[:size]
}

func Part1(input string) (string, error) {
	return string(checksum(fillDisk([]byte(input), diskSize))), nil
}

func Part2(input string) (string, error) {
	diskSize = 35651584
	return string(checksum(fillDisk([]byte(input), diskSize))), nil
}

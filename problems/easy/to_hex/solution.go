package tohex

var hexChars = []byte("0123456789abcdef")

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num += 1 << 32
	}

	var result []byte
	for num != 0 {
		result = append(result, hexChars[num&0xf])
		num >>= 4
	}

	return string(reverse(result))
}

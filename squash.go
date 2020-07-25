package squash

import "unicode"

func Space(s []byte) []byte {
	i := 0
	for range s {
		if (i + 1) == len(s) {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			continue
		}
		i++
	}
	return s
}

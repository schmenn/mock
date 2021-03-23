package mock

import (
	"unicode"
)

// Mock uppercases every second letter
func Mock(s string) string {
	out := []rune(s)
	for i := range out {
		if i%2 != 0 {
			out[i] = unicode.ToUpper(out[i])
			continue
		}
		out[i] = unicode.ToLower(out[i])
	}
	return string(out)
}

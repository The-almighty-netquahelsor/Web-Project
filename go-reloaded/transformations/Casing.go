package transformations

import (
	"strings"
	"unicode"
)

func applyCasing(mode, word string) string {
	switch mode {
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		r := []rune(word)
		if len(r) == 0 {
			return string(unicode.ToUpper(r[0])) + strings.ToLower(string(r[1:]))
		}
	}
	return word
}

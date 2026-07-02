package transformations

import "strings"

func FixArticles(word []string) []string {
	for j := 0; j < len(word)-1; j++ {
		if strings.EqualFold(word[j], "a") {
			next := strings.ToLower(strings.Trim(word[j+1], ".,;:?\"'"))
			if len(next) > 0 && strings.ContainsRune("aeiou", rune(next[0])) {
				if word[j] == "An" {
					word[j] = "A"
				} else {
					word[j] = "a"
				}
			}

		}
		if strings.EqualFold(word[j], "an") {
			next := strings.ToLower(strings.Trim(word[j+1], ".,;:?\"'"))
			if len(next) > 0 && strings.ContainsRune("aeiou", rune(next[0])) {
				if word[j] == "A" {
					word[j] = "An"
				} else {
					word[j] = "an"
				}
			}
		}
	}
	return word
}

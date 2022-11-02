package splitz

import (
	"unicode"
	"unicode/utf8"
)

func Splitz(src string) (result []string) {
	if !utf8.ValidString(src) {
		return []string{src}
	}

	result = []string{}
	var chars [][]rune
	lastClass := 0
	class := 0
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}

		if class == lastClass {
			chars[len(chars)-1] = append(chars[len(chars)-1], r)
		} else {
			chars = append(chars, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(chars)-1; i++ {
		if unicode.IsUpper(chars[i][0]) && unicode.IsLower(chars[i+1][0]) {
			chars[i+1] = append([]rune{
				chars[i][len(chars[i])-1],
			}, chars[i+1]...)
			chars[i] = chars[i][:len(chars[i])-1]
		}
	}

	for _, s := range chars {
		if len(s) > 0 {
			result = append(result, string(s))
		}
	}

	return result
}

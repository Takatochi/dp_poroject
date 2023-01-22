package stringFMT

import (
	"strings"
)

func StringTitleJoin(originalString string) string {
	words := strings.Fields(originalString)
	var result string
	for _, word := range words {
		result += strings.Title(word)
	}
	return result
}

package reverse_string

import "strings"

func ReverseString(input string) string {
	var output strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output.WriteRune(runes[i])
	}
	return output.String()
}

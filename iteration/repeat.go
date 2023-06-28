package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func BuiltInRepeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}

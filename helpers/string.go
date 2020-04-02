package helpers

import (
	"regexp"
	"strings"
)

// StringPointers converts a slice of string values into a slice of string pointers
//
// This function complements aws.StringSlice but works with variadic arguments so that an array literal is not required.
func StringPointers(strings ...string) []*string {
	sp := make([]*string, len(strings))
	for i := range sp {
		sp[i] = &strings[i]
	}
	return sp
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

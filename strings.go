package util

import (
	"regexp"
	"strings"
)

// CassandraFieldMapFromStruct ...
// func CassandraFieldMapFromStruct(st interface{}) []interface{} {

// }

// JoinStrings ... Joins an array of string
func JoinStrings(values ...string) string {
	return strings.Join(values, "")
}

// BytesToSize ...
// func BytesToSize(bytes float64) string {
// 	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}
// 	if bytes == 0 {
// 		return "0 Byte"
// 	}
// 	i := int(math.Floor(math.Log(bytes) / math.Log(1024)))

// 	return strconv.FormatFloat(math.Round((bytes / (math.Pow(1024, i) / 100))))

// 	// return string(i)
// }

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts camelCase to snake_case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ToCamel converts string to Upper CamelCase
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts string to lowerCamelCase
func ToLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func stringListToLowerCase(arr []string) []string {
	s := []string{}
	for _, field := range arr {
		s = append(s, ToSnakeCase(field))
	}

	return s
}

// RemoveTrailingWhiteSpaces removes all trailing whitespaces
func RemoveTrailingWhiteSpaces(text string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(text, " ")
	// return strings.Replace(text, "  ", "", -1)
}

// ToTitleCase removes all trailing whitespaces
func ToTitleCase(input string) string {
	input = strings.ToLower(input)
	words := strings.Fields(input)
	smallwords := " a an on to is of and in for "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

// StringContainsList checks whether a slice of string is found for a specified string
func StringContainsList(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}

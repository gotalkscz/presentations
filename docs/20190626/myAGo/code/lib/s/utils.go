package s

import (
	"regexp"
	"unicode"
)

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func D(id DepId) Dep {
	return Dep{
		Id: id,
	}
}

func Deps(ids ...DepId) (result []Dep) {
	for _, id := range ids {
		result = append(result, Dep{
			Id: id,
		})
	}
	return
}

// Make a Regex to say we only want letters and numbers
var reg = regexp.MustCompilePOSIX("[^a-zA-Z0-9]+")

func toName(in string) string {
	return ucFirst(reg.ReplaceAllString(in, ""))
}

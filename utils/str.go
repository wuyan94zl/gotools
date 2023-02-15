package utils

import (
	"strings"
)

func GetCamelCaseName(str string) string {
	str = strings.ReplaceAll(str, "_", "/")
	s := strings.Split(str, "/")
	rlt := ""
	for _, v := range s {
		rlt += strings.ToUpper(v[0:1]) + v[1:]
	}
	return rlt
}

func UpperFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

func ToLowers(str ...*string) {
	for i, v := range str {
		*str[i] = strings.ToLower(*v)
	}
}

func ToUppers(str ...*string) {
	for i, v := range str {
		*str[i] = strings.ToUpper(*v)
	}
}

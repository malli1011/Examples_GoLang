package codepractice

import "strings"

func Cat(str []string, sep string) string {
	s := str[0]
	for _, i := range str[1:] {
		s += sep
		s += i
	}
	return s
}

func Join(str []string, sep string) string {
	return strings.Join(str, sep)
}

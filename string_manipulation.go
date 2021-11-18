package strman

import (
	"regexp"
	"strings"
)

var (
	snakeKebabSplitter = regexp.MustCompile("[-_]")
	camelReplacer      = regexp.MustCompile("([a-z])([A-Z])")
)

const (
	kebabDelimiter = "-"
	snakeDelimiter = "_"
)

func split(source string) []string {
	source = camelReplacer.ReplaceAllString(source, "$1-$2")
	source = strings.ToLower(source)
	return snakeKebabSplitter.Split(source, -1)
}

func ToDelimited(source, delimiter string) string {
	return strings.Join(split(source), delimiter)
}

func ToScreamingDelimited(source, delimiter string) string {
	s := split(source)
	transform(s, strings.ToUpper, 0)
	return strings.Join(s, delimiter)
}

func ToKebab(source string) string {
	return ToDelimited(source, kebabDelimiter)
}

func ToScreamingKebab(source string) string {
	return ToScreamingDelimited(source, kebabDelimiter)
}

func ToSnake(source string) string {
	return ToDelimited(source, snakeDelimiter)
}

func ToScreamingSnake(source string) string {
	return ToScreamingDelimited(source, snakeDelimiter)
}

func ToCamel(source string) string {
	s := split(source)
	transform(s, strings.Title, 1)
	return strings.Join(s, "")
}

func ToPascal(source string) string {
	s := split(source)
	transform(s, strings.Title, 0)
	return strings.Join(s, "")
}

type transformFunc func(string) string

func transform(s []string, f transformFunc, n int) {
	for i, v := range s {
		if i < n {
			continue
		}
		s[i] = f(v)
	}
}

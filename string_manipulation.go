package strman

import (
	"strings"
)

const (
	kebabDelimiter = "-"
	snakeDelimiter = "_"
)

func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func isNum(b byte) bool {
	return '0' <= b && b <= '9'
}

func isSymbol(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r', '.', '_', '-':
		return false
	default:
		return true
	}
}

func Split(source string) []string {
	p := splitter{src: source}
	out := make([]string, 0, len(source)/3)
	for {
		next, done := p.next()
		if done {
			return out
		}
		out = append(out, next)
	}
}

func ToDelimited(source, delimiter string) string {
	return strings.Join(Split(source), delimiter)
}

func ToScreamingDelimited(source, delimiter string) string {
	s := Split(source)
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
	s := Split(source)
	transform(s, title, 1)
	return strings.Join(s, "")
}

func ToPascal(source string) string {
	s := Split(source)
	transform(s, title, 0)
	return strings.Join(s, "")
}

func title(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

func transform(src []string, transformFunc func(string) string, startIdx int) {
	for i, v := range src {
		if i < startIdx {
			continue
		}
		src[i] = transformFunc(v)
	}
}

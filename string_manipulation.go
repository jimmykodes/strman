package strman

import (
	"bytes"
	"strings"
)

const (
	kebabDelimiter = "-"
	snakeDelimiter = "_"
)

type lexer struct {
	source string
	pos    int
}

func (l *lexer) Peek() (byte, bool) {
	if l.pos >= len(l.source) {
		return 0, true
	}
	return l.source[l.pos], false
}

func (l *lexer) Read() (byte, bool) {
	b, done := l.Peek()
	if !done {
		l.pos++
	}
	return b, done
}

type parser struct {
	lexer  lexer
	buf    [128]byte
	bufLen int
	done   bool
}

func (p *parser) addByte(b byte) {
	p.buf[p.bufLen] = b
	p.bufLen++
}

func (p *parser) Next() ([]byte, bool) {
	for {
		peek, done := p.lexer.Peek()
		if done {
			p.done = done
			break
		}
		if peek == '-' || peek == '_' {
			// for sure a boundary, eat the divider and break
			p.lexer.Read()
			break
		}
		if isCapital(peek) {
			if p.bufLen > 0 && !isCapital(p.buf[p.bufLen-1]) {
				// the buffer isn't empty, and it isn't full of capitals, this is a word division
				break
			}
			p.buf[p.bufLen] = peek
			p.bufLen++
			p.lexer.Read()
		} else if isNum(peek) {
			if p.bufLen > 0 && !isNum(p.buf[p.bufLen-1]) {
				// the buffer isn't empty, and it isn't full of numbers, this is a word division
				break
			}
			p.buf[p.bufLen] = peek
			p.bufLen++
			p.lexer.Read()
		} else {
			p.buf[p.bufLen] = peek
			p.bufLen++
			p.lexer.Read()
		}
	}
	v := bytes.ToLower(p.buf[:p.bufLen])
	p.bufLen = 0
	return v, p.done
}

func isCapital(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func isNum(b byte) bool {
	return '0' <= b && b <= '9'
}

func split(source string) []string {
	l := lexer{source: source}
	p := parser{lexer: l}
	var out []string
	for {
		next, done := p.Next()
		out = append(out, string(next))
		if done {
			return out
		}
	}
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

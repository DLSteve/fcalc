package parser

import "bytes"

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	OPERATOR
	WHOLENUMBER
	NUMERATOR
	DENOMINATOR
	DIVIDER   // /
	SEPARATOR // _
	QUERY     // ?
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isNumeric(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isFractional(ch rune) bool {
	return isNumeric(ch) || ch == '-' || ch == '/' || ch == '_'
}

type Scanner struct {
	rs RuneStack
}

func NewScanner(s string) *Scanner {
	return &Scanner{
		rs: NewRuneStack(s),
	}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.rs.Pop()

	if isWhitespace(ch) {
		s.rs.Push(ch)
		return s.scanWhitespace()
	}

	if isOperator(ch) && isWhitespace(s.rs.Peek()) {
		return OPERATOR, string(ch)
	}

	if isFractional(ch) {
		s.rs.Push(ch)
		return s.scanFraction()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '?':
		return QUERY, string(ch)
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.rs.Pop())

	for {
		if ch := s.rs.Pop(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.rs.Push(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanFraction() (tok Token, lit string) {
	var buf bytes.Buffer

	for {
		ch := s.rs.Pop()
		next := s.rs.Peek()
		if ch == eof || isWhitespace(ch) {
			s.rs.Push(ch)
			break
		} else if isNumeric(ch) && next == '_' {
			buf.WriteRune(ch)
			return WHOLENUMBER, buf.String()
		} else if isNumeric(ch) && next == '/' {
			buf.WriteRune(ch)
			return NUMERATOR, buf.String()
		} else if isNumeric(ch) && (isWhitespace(next) || next == eof) {
			buf.WriteRune(ch)
			return DENOMINATOR, buf.String()
		} else if ch == '_' {
			buf.WriteRune(ch)
			return SEPARATOR, buf.String()
		} else if ch == '/' {
			buf.WriteRune(ch)
			return DIVIDER, buf.String()
		} else if isNumeric(ch) || (ch == '-' && isNumeric(next)) {
			buf.WriteRune(ch)
		} else {
			buf.WriteRune(ch)
			return ILLEGAL, buf.String()
		}
	}

	return ILLEGAL, buf.String()
}

package gopyparser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	EOF = rune(0)
)

type Scanner struct {
	r *bufio.Reader
}

func New(r io.Reader) *Scanner {
	s := &Scanner{
		r: bufio.NewReader(r),
	}
	return s
}

func NewFromString(s string) *Scanner {
	return New(strings.NewReader(s))
}

func NewFromFile(filePath string) (*Scanner, error) {

	r, e := os.Open(filePath)
	if e != nil {
		return nil, e
	}
	return New(r), nil
}

func (s *Scanner) Read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return EOF
	}
	return ch
}

func (s *Scanner) Unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) NextToken() Token {

	var tok Token

	ch := s.Read()
	switch ch {
	case '=':
		tok = NewToken(Equal, "=")
	case '(':
		tok = NewToken(LPar, "(")
	case ')':
		tok = NewToken(RPar, ")")
	default:
		if IsLetter(ch) {

		}
		tok = NewToken(ErrorToken, string(ch))
	}

	return tok
}

func IsSpace(ch rune) bool {
	return ch == ' '
}

func IsTab(ch rune) bool {
	return ch == '\t'
}

func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func IsLineBreak(c rune) bool {
	return c == '\n' || c == '\r'
}

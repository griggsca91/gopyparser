package gopyparser

import (
	"unicode/utf8"
	"unicode"
	"io/ioutil"
	"errors"
)

type Scanner struct {
	sourceText string

	lineIndex   int
	columnIndex int
}

func New(sourceText string) *Scanner {
	s := &Scanner{
		sourceText:  sourceText,
		lineIndex:   1,
		columnIndex: 1,
	}
	return s
}

func NewFromFile(filePath string) (*Scanner, error) {

	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}
	content := string(dat)

	return New(content), nil
}

func (s *Scanner) peekRune() rune {
	r, _ := utf8.DecodeRuneInString(s.sourceText)
	return r 
}



func IsSpace(c rune) bool {
	return uint32(c) == ' '
}

func IsTab(c rune) bool {
	return uint32(c) == '\t'
}

func IsLetter(c rune) bool {
	return unicode.IsLetter(c)
}

func IsLineBreak(c rune) bool {
	b := uint32(c)
	return b == '\n' || b == '\r' 
}
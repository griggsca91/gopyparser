package gopyparser

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

func NewFromFile(filePath string) *Scanner {

}

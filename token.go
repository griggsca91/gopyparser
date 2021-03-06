package gopyparser

const (
	EndMarker        = "EndMarker"
	Name             = "Name"
	Number           = "Number"
	String           = "String"
	NewLine          = "NewLine"
	Indent           = "Indent"
	Dedent           = "Dedent"
	LPar             = "LPar"
	RPar             = "RPar"
	LSQB             = "LSQB"
	RSQB             = "RSQB"
	Colon            = "Colon"
	Semi             = "Semi"
	Plus             = "Plus"
	Minus            = "Minus"
	Star             = "Star"
	Slash            = "Slash"
	VBar             = "VBar"
	Amper            = "Amper"
	Less             = "Less"
	Greater          = "Greater"
	Equal            = "Equal"
	Dot              = "Dot"
	Percent          = "Percent"
	LBrace           = "LBrace"
	RBrace           = "RBrace"
	EqEqual          = "EqEqual"
	NotEqual         = "NotEqual"
	LessEqual        = "LessEqual"
	GreaterEqual     = "GreaterEqual"
	Tilde            = "Tilde"
	Circumflex       = "Circumflex"
	LeftShift        = "LeftShift"
	RightShift       = "RightShift"
	DoubleStar       = "DoubleStar"
	PlusEqual        = "PlusEqual"
	MinEqual         = "MinEqual"
	StarEqual        = "StarEqual"
	SlashEqual       = "SlashEqual"
	PercentEqual     = "PercentEqual"
	AmperEqual       = "AmperEqual"
	VBarEqual        = "VBarEqual"
	CircumflexEqual  = "CircumflexEqual"
	LeftShiftEqual   = "LeftShiftEqual"
	RightShiftEqual  = "RightShiftEqual"
	DoubleStarEqual  = "DoubleStarEqual"
	DoubleSlash      = "DoubleSlash"
	DoubleSlashEqual = "DoubleSlashEqual"
	At               = "At"
	AtEqual          = "AtEqual"
	RArrow           = "RArrow"
	Ellipsis         = "Ellipsis"
	Op               = "Op"
	Await            = "Await"
	Async            = "Async"
	ErrorToken       = "<ErrorToken>"
	Comment          = "Comment"
	NL               = "NL"
	Encoding         = "Encoding"
	N_Tokens         = "<N_Tokens>"
)

type Token struct {
	LineIndex   int
	ColumnIndex int

	TokenType string
	Value     string
}

func NewToken(tokenType, value string) Token {
	return Token{
		TokenType: tokenType,
		Value:     value,
	}
}

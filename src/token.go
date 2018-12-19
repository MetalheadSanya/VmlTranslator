package src

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	NEW_LINE

	IMPORT
	AS

	IDENTIFIER
	STRING
	DOUBLE
	INTEGER

	DOT
	COLON
	SEMICOLON
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isNewLine(ch rune) bool {
	return ch == '\n'
}

func isDot(ch rune) bool {
	return ch == '.'
}

func isQuotationMark(ch rune) bool {
	return ch == '"'
}

func isReverseSolidus(ch rune) bool {
	return ch == '\\'
}

var eof = rune(0)

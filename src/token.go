package main

type Token int

const (
	ILLEGAL Token = iota
	EOF
	NEW_LINE

	IMPORT
	AS
	PROPERTY
	ALIAS
	READONLY
	SIGNAL
	ENUM

	IDENTIFIER

	STRING_LITERAL
	DOUBLE_LITERAL
	INT_LITERAL

	BOOL_TYPE
	DOUBLE_TYPE
	REAL_TYPE
	INT_TYPE
	STRING_TYPE
	VAR_TYPE

	DOT
	COLON
	SEMICOLON
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET

	SOLIDUS
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

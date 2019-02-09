package lexer

type Token = int

const (
	ILLEGAL Token = iota
	Eof
	NewLine

	Import
	As
	Default
	Property
	Alias
	Readonly
	Signal
	Enum
	True
	False

	Identifier

	StringLiteral
	DoubleLiteral
	IntLiteral

	BoolType
	DoubleType
	RealType
	IntType
	StringType
	VarType
	ListType

	Dot
	Comma
	Colon
	Semicolon
	LeftParenthesis
	RightParenthesis
	LeftCurlyBracket
	RightCurlyBracket
	LessThanSign
	MoreThanSign
	LeftSquareBracket
	RightSquareBracket

	Solidus
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

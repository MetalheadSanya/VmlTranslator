package main

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type Scanner struct {
	r *bufio.Reader
}

func VmlScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		s.missWhitespace()
		return s.Scan()
	} else if isNewLine(ch) {
		s.unread()
		return s.scanNewLine()
	} else if isDot(ch) {
		return DOT, "."
	} else if unicode.IsLetter(ch) {
		s.unread()
		return s.scanIdentifier()
	} else if unicode.IsDigit(ch) {
		s.unread()
		return s.scanNumber()
	} else if isQuotationMark(ch) {
		s.unread()
		return s.scanString()
	}

	switch ch {
	case '.':
		return DOT, string(ch)
	case ',':
		return COMMA, ","
	case eof:
		return EOF, ""
	case '(':
		return LEFT_PARENTHESIS, "("
	case ')':
		return RIGHT_PARENTHESIS, ")"
	case '{':
		return LEFT_CURLY_BRACKET, "{"
	case '}':
		return RIGHT_CURLY_BRACKET, "}"
	case ';':
		return SEMICOLON, ";"
	case ':':
		return COLON, ":"
	case '/':
		s.unread()
		if s.missComment() {
			return s.Scan()
		} else {
			return SOLIDUS, "/"
		}
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) missWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}

func (s *Scanner) scanNewLine() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNewLine(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NEW_LINE, buf.String()
}

func (s *Scanner) scanIdentifier() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch buf.String() {
	case "import":
		return IMPORT, buf.String()
	case "as":
		return AS, buf.String()
	case "property":
		return PROPERTY, buf.String()
	case "alias":
		return ALIAS, buf.String()
	case "readonly":
		return READONLY, buf.String()
	case "signal":
		return SIGNAL, buf.String()
	case "enum":
		return ENUM, buf.String()
	case "int":
		return INT_TYPE, buf.String()
	case "bool":
		return BOOL_TYPE, buf.String()
	case "double":
		return DOUBLE_TYPE, buf.String()
	case "real":
		return REAL_TYPE, buf.String()
	case "string":
		return STRING_TYPE, buf.String()
	case "var":
		return VAR_TYPE, buf.String()
	}

	return IDENTIFIER, buf.String()
}

func (s *Scanner) scanNumber() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	wasDot := false
	wasDigit := false
	brokenToken := false

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDot(ch) && !unicode.IsDigit(ch) {
			s.unread()
			break
		} else {
			if isDot(ch) && !wasDot {
				wasDot = true
			} else if isDot(ch) {
				brokenToken = true
			} else if unicode.IsDigit(ch) && wasDot {
				wasDigit = true
			}
			_, _ = buf.WriteRune(ch)
		}
	}

	if wasDot && !wasDigit {
		brokenToken = true
	}

	if brokenToken {
		return ILLEGAL, buf.String()
	} else if wasDot {
		return DOUBLE_LITERAL, buf.String()
	} else {
		return INT_LITERAL, buf.String()
	}
}

func (s *Scanner) scanString() (tok Token, lit string) {
	var buf bytes.Buffer
	s.read()

	wasQuotationMark := false
	wasEscape := false

	for {
		if ch := s.read(); ch == eof {
			break
		} else {
			if isReverseSolidus(ch) {
				wasEscape = true
			} else if !wasEscape && isQuotationMark(ch) {
				wasQuotationMark = true
				break
			} else if isNewLine(ch) {
				s.unread()
				break
			} else {
				wasEscape = false
			}

			_, _ = buf.WriteRune(ch)
		}
	}

	if wasQuotationMark {
		return STRING_LITERAL, buf.String()
	} else {
		return ILLEGAL, buf.String()
	}
}

func (s *Scanner) missComment() bool {
	s.read()

	ch := s.read()

	switch ch {
	case '/':
		for {
			ch = s.read()
			if ch == eof {
				return true
			} else if ch == '\n' {
				s.unread()
				return true
			}
		}
	case '*':
		for {
			ch = s.read()
			if ch == eof {
				return true
			} else if ch == '*' {
				ch = s.read()
				if ch == '/' {
					return true
				}
				s.unread()
			}
		}
	}

	s.unread()

	return false
}

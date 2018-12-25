package lexer

import (
	"strings"
	"testing"
)

func TestLexerList(t *testing.T) {
	str := `Rectangle {
    property list<Rectangle> siblingRects
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != ListType {
		t.Errorf("%s is not ListType token", lex)
	}
	if token, lex := s.Scan(); token != LessThanSign {
		t.Errorf("%s is not LessThanSign token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != MoreThanSign {
		t.Errorf("%s is not MoreThanSign token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "siblingRects" {
		t.Errorf("%s is not Identifier(siblingRects) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerListWithValue(t *testing.T) {
	str := `Rectangle {
    property list<Rectangle> childRects: [
        Rectangle { color: "red" },
        Rectangle { color: "blue"}
    ]
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != ListType {
		t.Errorf("%s is not ListType token", lex)
	}
	if token, lex := s.Scan(); token != LessThanSign {
		t.Errorf("%s is not LessThanSign token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != MoreThanSign {
		t.Errorf("%s is not MoreThanSign token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "childRects" {
		t.Errorf("%s is not Identifier(childRects) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != LeftSquareBracket {
		t.Errorf("%s is not LeftSquareBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != StringLiteral || lex != "red" {
		t.Errorf("%s is not StringLiteral(red) token", lex)
	}
	if token, lex := s.Scan(); token != RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != Comma {
		t.Errorf("%s is not Comma token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != StringLiteral || lex != "blue" {
		t.Errorf("%s is not StringLiteral(blue) token", lex)
	}
	if token, lex := s.Scan(); token != RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != RightSquareBracket {
		t.Errorf("%s is not RightSquareBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

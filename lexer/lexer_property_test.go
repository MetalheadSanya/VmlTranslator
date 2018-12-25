package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
	"strings"
	"testing"
)

func TestLexerProperty(t *testing.T) {
	str := `Rectangle {
    property color previousColor
    property color nextColor
}`
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "previousColor" {
		t.Errorf("%s is not Identifier(previousColor) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "nextColor" {
		t.Errorf("%s is not Identifier(nextColor) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerPropertyWithValue(t *testing.T) {
	str := `Rectangle {
    property color nextColor: "blue"
}`
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "nextColor" {
		t.Errorf("%s is not Identifier(previousColor) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "blue" {
		t.Errorf("%s is not StringLiteral(blue) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerPropertyReadonly(t *testing.T) {
	str := `Rectangle {
    property alias color: rectangle.border.color
}`
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Alias {
		t.Errorf("%s is not Alias token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "rectangle" {
		t.Errorf("%s is not Identifier(rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Dot {
		t.Errorf("%s is not Dot token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "border" {
		t.Errorf("%s is not Identifier(border) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Dot {
		t.Errorf("%s is not Dot token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerPropertyAlias(t *testing.T) {
	str := `Rectangle {
    readonly property int someNumber: 10
}`
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Identifier(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Readonly {
		t.Errorf("%s is not Readonly token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.IntType {
		t.Errorf("%s is not IntType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "someNumber" {
		t.Errorf("%s is not Identifier(someNumber) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.IntLiteral || lex != "10" {
		t.Errorf("%s is not IntLiteral(10) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not RightCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

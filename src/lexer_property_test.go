package main

import (
	"strings"
	"testing"
)

func TestLexerProperty(t *testing.T) {
	str := `Rectangle {
    property color previousColor
    property color nextColor
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "previousColor" {
		t.Errorf("%s is not IDENTIFIER(previousColor) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "nextColor" {
		t.Errorf("%s is not IDENTIFIER(nextColor) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerPropertyWithValue(t *testing.T) {
	str := `Rectangle {
    property color nextColor: "blue"
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "nextColor" {
		t.Errorf("%s is not IDENTIFIER(previousColor) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "blue" {
		t.Errorf("%s is not STRING_LITERAL(blue) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerPropertyReadonly(t *testing.T) {
	str := `Rectangle {
    property alias color: rectangle.border.color
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != ALIAS {
		t.Errorf("%s is not ALIAS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "rectangle" {
		t.Errorf("%s is not IDENTIFIER(rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != DOT {
		t.Errorf("%s is not DOT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "border" {
		t.Errorf("%s is not IDENTIFIER(border) token", lex)
	}
	if token, lex := s.Scan(); token != DOT {
		t.Errorf("%s is not DOT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerPropertyAlias(t *testing.T) {
	str := `Rectangle {
    readonly property int someNumber: 10
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != READONLY {
		t.Errorf("%s is not READONLY token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != INT_TYPE {
		t.Errorf("%s is not INT_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "someNumber" {
		t.Errorf("%s is not IDENTIFIER(someNumber) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "10" {
		t.Errorf("%s is not INT_LITERAL(10) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

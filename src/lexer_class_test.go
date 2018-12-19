package src

import (
	"strings"
	"testing"
)

func TestLexerClass(t *testing.T) {
	str := `Rectangle {
    width: 100.0
    height: 10
    color: "red"
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE || lex != "100.0" {
		t.Errorf("%s is not INTEGER(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "height" {
		t.Errorf("%s is not IDENTIFIER(height) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != INTEGER || lex != "10" {
		t.Errorf("%s is not INTEGER(10) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != STRING || lex != "red" {
		t.Errorf("%s is not STRING(red) token", lex)
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

func TestLexerClassSingleLine(t *testing.T) {
	str := `Rectangle { width: 100.0; height: 10; color: "red" }`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE || lex != "100.0" {
		t.Errorf("%s is not DOUBLE(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != SEMICOLON {
		t.Errorf("%s is not SEMICOLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "height" {
		t.Errorf("%s is not IDENTIFIER(height) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != INTEGER || lex != "10" {
		t.Errorf("%s is not INTEGER(10) token", lex)
	}
	if token, lex := s.Scan(); token != SEMICOLON {
		t.Errorf("%s is not SEMICOLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != STRING || lex != "red" {
		t.Errorf("%s is not STRING(red) token", lex)
	}
	if token, lex := s.Scan(); token != WS {
		t.Errorf("%s is not WS token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

package lexer

import (
	"strings"
	"testing"
)

func TestLexerSingleLineComments(t *testing.T) {
	str := `Rectangle {
    width: 100.0// Test
    height: 10 //--// A-ha-ha
    color: "red"
}/////////`
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
	if token, lex := s.Scan(); token != Identifier || lex != "width" {
		t.Errorf("%s is not Identifier(width) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != DoubleLiteral || lex != "100.0" {
		t.Errorf("%s is not DoubleLiteral(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "height" {
		t.Errorf("%s is not Identifier(height) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != IntLiteral || lex != "10" {
		t.Errorf("%s is not IntLiteral(10) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
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

func TestLexerMultilineComments(t *testing.T) {
	str := `Rectangle {
    width: 100.0/* Test
    height: 10 */
    color: "red"/**/
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
	if token, lex := s.Scan(); token != Identifier || lex != "width" {
		t.Errorf("%s is not Identifier(width) token", lex)
	}
	if token, lex := s.Scan(); token != Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != DoubleLiteral || lex != "100.0" {
		t.Errorf("%s is not DoubleLiteral(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
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

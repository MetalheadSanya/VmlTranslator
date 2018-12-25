package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
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
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "width" {
		t.Errorf("%s is not Identifier(width) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "100.0" {
		t.Errorf("%s is not DoubleLiteral(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "height" {
		t.Errorf("%s is not Identifier(height) token", lex)
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
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "red" {
		t.Errorf("%s is not StringLiteral(red) token", lex)
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

func TestLexerMultilineComments(t *testing.T) {
	str := `Rectangle {
    width: 100.0/* Test
    height: 10 */
    color: "red"/**/
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
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "width" {
		t.Errorf("%s is not Identifier(width) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "100.0" {
		t.Errorf("%s is not DoubleLiteral(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "red" {
		t.Errorf("%s is not StringLiteral(red) token", lex)
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

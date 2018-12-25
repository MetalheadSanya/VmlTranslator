package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
	"strings"
	"testing"
)

func TestLexerTypes(t *testing.T) {
	str := `Rectangle {
    property int intProperty
    property string stringProperty
	property bool boolProperty
	property double doubleProperty
	property real realProperty
	property var varProperty
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
	if token, lex := s.Scan(); token != lexer.IntType {
		t.Errorf("%s is not IntType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "intProperty" {
		t.Errorf("%s is not Identifier(intProperty) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringType {
		t.Errorf("%s is not StringType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "stringProperty" {
		t.Errorf("%s is not Identifier(stringProperty) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.BoolType {
		t.Errorf("%s is not BoolType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "boolProperty" {
		t.Errorf("%s is not Identifier(boolProperty) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleType {
		t.Errorf("%s is not DoubleType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "doubleProperty" {
		t.Errorf("%s is not Identifier(doubleProperty) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RealType {
		t.Errorf("%s is not RealType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "realProperty" {
		t.Errorf("%s is not Identifier(realProperty) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != lexer.VarType {
		t.Errorf("%s is not VarType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "varProperty" {
		t.Errorf("%s is not Identifier(varProperty) token", lex)
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

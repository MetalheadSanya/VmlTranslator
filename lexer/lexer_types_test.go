package lexer

import (
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
	if token, lex := s.Scan(); token != IntType {
		t.Errorf("%s is not IntType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "intProperty" {
		t.Errorf("%s is not Identifier(intProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != StringType {
		t.Errorf("%s is not StringType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "stringProperty" {
		t.Errorf("%s is not Identifier(stringProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != BoolType {
		t.Errorf("%s is not BoolType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "boolProperty" {
		t.Errorf("%s is not Identifier(boolProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != DoubleType {
		t.Errorf("%s is not DoubleType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "doubleProperty" {
		t.Errorf("%s is not Identifier(doubleProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != RealType {
		t.Errorf("%s is not RealType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "realProperty" {
		t.Errorf("%s is not Identifier(realProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}
	if token, lex := s.Scan(); token != Property {
		t.Errorf("%s is not Property token", lex)
	}
	if token, lex := s.Scan(); token != VarType {
		t.Errorf("%s is not VarType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "varProperty" {
		t.Errorf("%s is not Identifier(varProperty) token", lex)
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

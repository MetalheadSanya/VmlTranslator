package main

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
	if token, lex := s.Scan(); token != INT_TYPE {
		t.Errorf("%s is not INT_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "intProperty" {
		t.Errorf("%s is not IDENTIFIER(intProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != STRING_TYPE {
		t.Errorf("%s is not STRING_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "stringProperty" {
		t.Errorf("%s is not IDENTIFIER(stringProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != BOOL_TYPE {
		t.Errorf("%s is not BOOL_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "boolProperty" {
		t.Errorf("%s is not IDENTIFIER(boolProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE_TYPE {
		t.Errorf("%s is not DOUBLE_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "doubleProperty" {
		t.Errorf("%s is not IDENTIFIER(doubleProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != REAL_TYPE {
		t.Errorf("%s is not REAL_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "realProperty" {
		t.Errorf("%s is not IDENTIFIER(realProperty) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != PROPERTY {
		t.Errorf("%s is not PROPERTY token", lex)
	}
	if token, lex := s.Scan(); token != VAR_TYPE {
		t.Errorf("%s is not VAR_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "varProperty" {
		t.Errorf("%s is not IDENTIFIER(varProperty) token", lex)
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

package main

import (
	"strings"
	"testing"
)

func TestLexerSignal(t *testing.T) {
	str := `Item {
    signal clicked
    signal hovered()
    signal actionPerformed(string action, var actionResult)
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Identifier || lex != "Item" {
		t.Errorf("%s is not Identifier(Item) token", lex)
	}
	if token, lex := s.Scan(); token != LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "clicked" {
		t.Errorf("%s is not Identifier(clicked) token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "hovered" {
		t.Errorf("%s is not Identifier(hovered) token", lex)
	}
	if token, lex := s.Scan(); token != LeftParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != RightParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "actionPerformed" {
		t.Errorf("%s is not Identifier(actionPerformed) token", lex)
	}
	if token, lex := s.Scan(); token != LeftParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != StringType {
		t.Errorf("%s is not StringType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "action" {
		t.Errorf("%s is not Identifier(action) token", lex)
	}
	if token, lex := s.Scan(); token != Comma {
		t.Errorf("%s is not Comma token", lex)
	}
	if token, lex := s.Scan(); token != VarType {
		t.Errorf("%s is not VarType token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "actionResult" {
		t.Errorf("%s is not Identifier(actionResult) token", lex)
	}
	if token, lex := s.Scan(); token != RightParenthesis {
		t.Errorf("%s is not Signal token", lex)
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

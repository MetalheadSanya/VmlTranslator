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
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Item" {
		t.Errorf("%s is not IDENTIFIER(Item) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != SIGNAL {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "clicked" {
		t.Errorf("%s is not IDENTIFIER(clicked) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != SIGNAL {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "hovered" {
		t.Errorf("%s is not IDENTIFIER(hovered) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_PARENTHESIS {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_PARENTHESIS {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != SIGNAL {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "actionPerformed" {
		t.Errorf("%s is not IDENTIFIER(actionPerformed) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_PARENTHESIS {
		t.Errorf("%s is not SIGNAL token", lex)
	}
	if token, lex := s.Scan(); token != STRING_TYPE {
		t.Errorf("%s is not STRING_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "action" {
		t.Errorf("%s is not IDENTIFIER(action) token", lex)
	}
	if token, lex := s.Scan(); token != COMMA {
		t.Errorf("%s is not COMMA token", lex)
	}
	if token, lex := s.Scan(); token != VAR_TYPE {
		t.Errorf("%s is not VAR_TYPE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "actionResult" {
		t.Errorf("%s is not IDENTIFIER(actionResult) token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_PARENTHESIS {
		t.Errorf("%s is not SIGNAL token", lex)
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

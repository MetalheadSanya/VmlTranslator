package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
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
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Item" {
		t.Errorf("%s is not Identifier(Item) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "clicked" {
		t.Errorf("%s is not Identifier(clicked) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "hovered" {
		t.Errorf("%s is not Identifier(hovered) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Signal {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "actionPerformed" {
		t.Errorf("%s is not Identifier(actionPerformed) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftParenthesis {
		t.Errorf("%s is not Signal token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringType {
		t.Errorf("%s is not StringType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "action" {
		t.Errorf("%s is not Identifier(action) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Comma {
		t.Errorf("%s is not Comma token", lex)
	}
	if token, lex := s.Scan(); token != lexer.VarType {
		t.Errorf("%s is not VarType token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "actionResult" {
		t.Errorf("%s is not Identifier(actionResult) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightParenthesis {
		t.Errorf("%s is not Signal token", lex)
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

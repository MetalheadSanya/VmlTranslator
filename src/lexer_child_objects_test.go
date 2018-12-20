package main

import (
	"strings"
	"testing"
)

func TestLexerChildObjectInProperty(t *testing.T) {
	str := `Rectangle {
    width: 100
    height: 100

    gradient: Gradient {
        GradientStop { position: 0.0; color: "yellow" }
        GradientStop { position: 1.0; color: "green" }
    }
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "100" {
		t.Errorf("%s is not INT_LITERAL(100) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "height" {
		t.Errorf("%s is not IDENTIFIER(height) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "100" {
		t.Errorf("%s is not INT_LITERAL(100) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "gradient" {
		t.Errorf("%s is not IDENTIFIER(gradient) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Gradient" {
		t.Errorf("%s is not IDENTIFIER(Gradient) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "GradientStop" {
		t.Errorf("%s is not IDENTIFIER(GradientStop) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "position" {
		t.Errorf("%s is not IDENTIFIER(position) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE_LITERAL || lex != "0.0" {
		t.Errorf("%s is not DOUBLE_LITERAL(0.0) token", lex)
	}
	if token, lex := s.Scan(); token != SEMICOLON {
		t.Errorf("%s is not SEMICOLON token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "yellow" {
		t.Errorf("%s is not STRING_LITERAL(yellow) token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "GradientStop" {
		t.Errorf("%s is not IDENTIFIER(GradientStop) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "position" {
		t.Errorf("%s is not IDENTIFIER(position) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE_LITERAL || lex != "1.0" {
		t.Errorf("%s is not DOUBLE_LITERAL(1.0) token", lex)
	}
	if token, lex := s.Scan(); token != SEMICOLON {
		t.Errorf("%s is not SEMICOLON token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "green" {
		t.Errorf("%s is not STRING_LITERAL(green) token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerChildObject(t *testing.T) {
	str := `Rectangle {
    width: 200
    height: 200
    color: "red"

    Text {
        anchors.centerIn: parent
        text: "Hello, QML!"
    }
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "200" {
		t.Errorf("%s is not INT_LITERAL(200) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "height" {
		t.Errorf("%s is not IDENTIFIER(height) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "200" {
		t.Errorf("%s is not INT_LITERAL(200) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "red" {
		t.Errorf("%s is not STRING_LITERAL(red) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Text" {
		t.Errorf("%s is not IDENTIFIER(Text) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "anchors" {
		t.Errorf("%s is not IDENTIFIER(anchors) token", lex)
	}
	if token, lex := s.Scan(); token != DOT {
		t.Errorf("%s is not DOT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "centerIn" {
		t.Errorf("%s is not IDENTIFIER(centerIn) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "parent" {
		t.Errorf("%s is not IDENTIFIER(parent) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != IDENTIFIER || lex != "text" {
		t.Errorf("%s is not IDENTIFIER(text) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "Hello, QML!" {
		t.Errorf("%s is not STRING_LITERAL(Hello, QML!) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}

	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF {
		t.Errorf("%s is not EOF token", lex)
	}
}

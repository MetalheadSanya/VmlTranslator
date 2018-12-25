package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
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
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Import token", lex)
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
	if token, lex := s.Scan(); token != lexer.IntLiteral || lex != "100" {
		t.Errorf("%s is not IntLiteral(100) token", lex)
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
	if token, lex := s.Scan(); token != lexer.IntLiteral || lex != "100" {
		t.Errorf("%s is not IntLiteral(100) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "gradient" {
		t.Errorf("%s is not Identifier(gradient) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Gradient" {
		t.Errorf("%s is not Identifier(Gradient) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "GradientStop" {
		t.Errorf("%s is not Identifier(GradientStop) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "position" {
		t.Errorf("%s is not Identifier(position) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "0.0" {
		t.Errorf("%s is not DoubleLiteral(0.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Semicolon {
		t.Errorf("%s is not Semicolon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "yellow" {
		t.Errorf("%s is not StringLiteral(yellow) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "GradientStop" {
		t.Errorf("%s is not Identifier(GradientStop) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "position" {
		t.Errorf("%s is not Identifier(position) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "1.0" {
		t.Errorf("%s is not DoubleLiteral(1.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Semicolon {
		t.Errorf("%s is not Semicolon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "color" {
		t.Errorf("%s is not Identifier(color) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "green" {
		t.Errorf("%s is not StringLiteral(green) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof {
		t.Errorf("%s is not Eof token", lex)
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
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Rectangle" {
		t.Errorf("%s is not Import token", lex)
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
	if token, lex := s.Scan(); token != lexer.IntLiteral || lex != "200" {
		t.Errorf("%s is not IntLiteral(200) token", lex)
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
	if token, lex := s.Scan(); token != lexer.IntLiteral || lex != "200" {
		t.Errorf("%s is not IntLiteral(200) token", lex)
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

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Text" {
		t.Errorf("%s is not Identifier(Text) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.LeftCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "anchors" {
		t.Errorf("%s is not Identifier(anchors) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Dot {
		t.Errorf("%s is not Dot token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "centerIn" {
		t.Errorf("%s is not Identifier(centerIn) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "parent" {
		t.Errorf("%s is not Identifier(parent) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.Identifier || lex != "text" {
		t.Errorf("%s is not Identifier(text) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Colon {
		t.Errorf("%s is not Colon token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "Hello, QML!" {
		t.Errorf("%s is not StringLiteral(Hello, QML!) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.NewLine {
		t.Errorf("%s is not NewLine token", lex)
	}

	if token, lex := s.Scan(); token != lexer.RightCurlyBracket {
		t.Errorf("%s is not LeftCurlyBracket token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof {
		t.Errorf("%s is not Eof token", lex)
	}
}

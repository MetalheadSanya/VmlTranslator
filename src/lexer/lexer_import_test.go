package lexer

import (
	"strings"
	"testing"
)

func TestLexerImportModule(t *testing.T) {
	str := "import QtQuick 2.0"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "QtQuick" {
		t.Errorf("%s is not Identifier token", lex)
	}
	if token, lex := s.Scan(); token != DoubleLiteral || lex != "2.0" {
		t.Errorf("%s is not DoubleLiteral(2.0) token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportModuleFull(t *testing.T) {
	str := "import QtQuick.LocalStorage 2.0 as Database"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "QtQuick" {
		t.Errorf("%s is not Identifier(QtQuick) token", lex)
	}
	if token, lex := s.Scan(); token != Dot || lex != "." {
		t.Errorf("%s is not Dot token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "LocalStorage" {
		t.Errorf("%s is not Identifier(LocalStorage) token", lex)
	}
	if token, lex := s.Scan(); token != DoubleLiteral || lex != "2.0" {
		t.Errorf("%s is not DoubleLiteral(2.0) token", lex)
	}
	if token, lex := s.Scan(); token != As || lex != "as" {
		t.Errorf("%s is not As token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "Database" {
		t.Errorf("%s is not Identifier(Database) token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportDirectory(t *testing.T) {
	str := "import \"directory\""
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != StringLiteral || lex != "directory" {
		t.Errorf("%s is not StringLiteral(directory) token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportDirectoryFull(t *testing.T) {
	str := "import \"file.js\" as ScriptIdentifier"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != StringLiteral || lex != "file.js" {
		t.Errorf("%s is not StringLiteral(file.js) token", lex)
	}
	if token, lex := s.Scan(); token != As || lex != "as" {
		t.Errorf("%s is not As token", lex)
	}
	if token, lex := s.Scan(); token != Identifier || lex != "ScriptIdentifier" {
		t.Errorf("%s is not Identifier(ScriptIdentifier) token", lex)
	}
	if token, lex := s.Scan(); token != Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

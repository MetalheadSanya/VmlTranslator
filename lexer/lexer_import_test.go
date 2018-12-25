package lexer_test

import (
	"github.com/MetalheadSanya/VmlTranslator/lexer"
	"strings"
	"testing"
)

func TestLexerImportModule(t *testing.T) {
	str := "import QtQuick 2.0"
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "QtQuick" {
		t.Errorf("%s is not Identifier token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "2.0" {
		t.Errorf("%s is not DoubleLiteral(2.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportModuleFull(t *testing.T) {
	str := "import QtQuick.LocalStorage 2.0 as Database"
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "QtQuick" {
		t.Errorf("%s is not Identifier(QtQuick) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Dot || lex != "." {
		t.Errorf("%s is not Dot token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "LocalStorage" {
		t.Errorf("%s is not Identifier(LocalStorage) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.DoubleLiteral || lex != "2.0" {
		t.Errorf("%s is not DoubleLiteral(2.0) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.As || lex != "as" {
		t.Errorf("%s is not As token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "Database" {
		t.Errorf("%s is not Identifier(Database) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportDirectory(t *testing.T) {
	str := "import \"directory\""
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "directory" {
		t.Errorf("%s is not StringLiteral(directory) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

func TestLexerImportDirectoryFull(t *testing.T) {
	str := "import \"file.js\" as ScriptIdentifier"
	r := strings.NewReader(str)
	s := lexer.VmlScanner(r)
	if token, lex := s.Scan(); token != lexer.Import || lex != "import" {
		t.Errorf("%s is not Import token", lex)
	}
	if token, lex := s.Scan(); token != lexer.StringLiteral || lex != "file.js" {
		t.Errorf("%s is not StringLiteral(file.js) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.As || lex != "as" {
		t.Errorf("%s is not As token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Identifier || lex != "ScriptIdentifier" {
		t.Errorf("%s is not Identifier(ScriptIdentifier) token", lex)
	}
	if token, lex := s.Scan(); token != lexer.Eof || lex != "" {
		t.Errorf("%s is not Eof token", lex)
	}
}

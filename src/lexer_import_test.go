package main

import (
	"strings"
	"testing"
)

func TestLexerImportModule(t *testing.T) {
	str := "import QtQuick 2.0"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IMPORT || lex != "import" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "QtQuick" {
		t.Errorf("%s is not IDENTIFIER token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE || lex != "2.0" {
		t.Errorf("%s is not DOUBLE(2.0) token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerImportModuleFull(t *testing.T) {
	str := "import QtQuick.LocalStorage 2.0 as Database"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IMPORT || lex != "import" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "QtQuick" {
		t.Errorf("%s is not IDENTIFIER(QtQuick) token", lex)
	}
	if token, lex := s.Scan(); token != DOT || lex != "." {
		t.Errorf("%s is not DOT token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "LocalStorage" {
		t.Errorf("%s is not IDENTIFIER(LocalStorage) token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE || lex != "2.0" {
		t.Errorf("%s is not DOUBLE token", lex)
	}
	if token, lex := s.Scan(); token != AS || lex != "as" {
		t.Errorf("%s is not AS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Database" {
		t.Errorf("%s is not IDENTIFIER(Database) token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerImportDirectory(t *testing.T) {
	str := "import \"directory\""
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IMPORT || lex != "import" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != STRING || lex != "directory" {
		t.Errorf("%s is not STRING(directory) token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerImportDirectoryFull(t *testing.T) {
	str := "import \"file.js\" as ScriptIdentifier"
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IMPORT || lex != "import" {
		t.Errorf("%s is not IMPORT token", lex)
	}
	if token, lex := s.Scan(); token != STRING || lex != "file.js" {
		t.Errorf("%s is not STRING(file.js) token", lex)
	}
	if token, lex := s.Scan(); token != AS || lex != "as" {
		t.Errorf("%s is not AS token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "ScriptIdentifier" {
		t.Errorf("%s is not IDENTIFIER(ScriptIdentifier) token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

package parser_test

import (
	"github.com/MetalheadSanya/VmlTranslator/parser"
	"strings"
	"testing"
)

func TestParserImportModule(t *testing.T) {
	str := "import QtQuick 2.0"
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	namespaceImportCount := len(stmt.NamespaceImports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.NamespaceImports[0]
	moduleIdentifierCount := len(importStmt.ModuleIdentifier)
	if moduleIdentifierCount != 1 {
		t.Errorf("Incorrect module identifier uri, expected 1, take %d", moduleIdentifierCount)
		return
	}
	if importStmt.ModuleIdentifier[0] != "QtQuick" {
		t.Errorf("Incorrect namespace, expected QtQuick, take %s", importStmt.ModuleIdentifier)
		return
	}
	if importStmt.Version != 2.0 {
		t.Errorf("Incorrect version, expected 2.0, take %f", importStmt.Version)
		return
	}
	if importStmt.Qualifier != nil {
		t.Error("Qualifier exist")
		return
	}
}

func TestParserImportModuleFull(t *testing.T) {
	str := "import QtQuick.LocalStorage 2.0 as Database"
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	namespaceImportCount := len(stmt.NamespaceImports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.NamespaceImports[0]
	moduleIdentifierCount := len(importStmt.ModuleIdentifier)
	if moduleIdentifierCount != 2 {
		t.Errorf("Incorrect module identifier uri, expected 2, take %d", moduleIdentifierCount)
		return
	}
	if importStmt.ModuleIdentifier[0] != "QtQuick" {
		t.Errorf("Incorrect namespace, expected NewReader, take %s", importStmt.ModuleIdentifier)
		return
	}
	if importStmt.ModuleIdentifier[1] != "LocalStorage" {
		t.Errorf("Incorrect namespace, expected NewReader, take %s", importStmt.ModuleIdentifier)
		return
	}
	if importStmt.Version != 2.0 {
		t.Errorf("Incorrect version, expected 2.0, take %f", importStmt.Version)
		return
	}
	if *importStmt.Qualifier != "Database" {
		t.Errorf("Incorrect qualifier, expected Database, take %q", *importStmt.Qualifier)
		return
	}
}

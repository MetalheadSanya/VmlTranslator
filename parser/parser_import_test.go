package parser_test

import (
	"github.com/MetalheadSanya/VmlTranslator/parser"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
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
	namespaceImportCount := len(stmt.Imports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.Imports[0].(*statement.ImportNamespace)
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
	namespaceImportCount := len(stmt.Imports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.Imports[0].(*statement.ImportNamespace)
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

func TestParserImportDirectory(t *testing.T) {
	str := "import \"UIKit\""
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	namespaceImportCount := len(stmt.Imports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.Imports[0].(*statement.ImportDirectory)
	if importStmt.Directory != "UIKit" {
		t.Errorf("Incorrect directory, expected UIKit, take %s", importStmt.Directory)
		return
	}
	if importStmt.Qualifier != nil {
		t.Error("Qualifier exist")
		return
	}
}

func TestParserImportDirectoryFull(t *testing.T) {
	str := "import \"UIKit\" as UI"
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	namespaceImportCount := len(stmt.Imports)
	if namespaceImportCount != 1 {
		t.Errorf("Incorrect import statements count, expected 1, take %d", namespaceImportCount)
		return
	}
	importStmt := stmt.Imports[0].(*statement.ImportDirectory)
	if importStmt.Directory != "UIKit" {
		t.Errorf("Incorrect directory, expected UIKit, take %s", importStmt.Directory)
		return
	}
	if *importStmt.Qualifier != "UI" {
		t.Errorf("Incorrect qualifier, expected UI, take %q", *importStmt.Qualifier)
		return
	}
}

func TestParserImports(t *testing.T) {
	str := `import "UIKit" as UI
import QtQuick 2.0 
`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	namespaceImportCount := len(stmt.Imports)
	if namespaceImportCount != 2 {
		t.Errorf("Incorrect import statements count, expected 2, take %d", namespaceImportCount)
		return
	}
	dirStmt := stmt.Imports[0].(*statement.ImportDirectory)
	if dirStmt.Directory != "UIKit" {
		t.Errorf("Incorrect directory, expected UIKit, take %s", dirStmt.Directory)
		return
	}
	if *dirStmt.Qualifier != "UI" {
		t.Errorf("Incorrect qualifier, expected UI, take %q", *dirStmt.Qualifier)
		return
	}
	namespaceStmt := stmt.Imports[1].(*statement.ImportNamespace)
	moduleIdentifierCount := len(namespaceStmt.ModuleIdentifier)
	if moduleIdentifierCount != 1 {
		t.Errorf("Incorrect module identifier uri, expected 1, take %d", moduleIdentifierCount)
		return
	}
	if namespaceStmt.ModuleIdentifier[0] != "QtQuick" {
		t.Errorf("Incorrect namespace, expected QtQuick, take %s", namespaceStmt.ModuleIdentifier)
		return
	}
	if namespaceStmt.Version != 2.0 {
		t.Errorf("Incorrect version, expected 2.0, take %f", namespaceStmt.Version)
		return
	}
	if namespaceStmt.Qualifier != nil {
		t.Error("Qualifier exist")
		return
	}
}

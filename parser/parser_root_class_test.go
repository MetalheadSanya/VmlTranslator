package parser_test

import (
	"github.com/MetalheadSanya/VmlTranslator/parser"
	"strings"
	"testing"
)

func TestParserRootClassSingleLine(t *testing.T) {
	str := "Rect { }"
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.ClassStatement
	if classStmt.Name != "Rect" {
		t.Errorf("Incorrect class name, expected Rect, take %q", classStmt.Name)
		return
	}
}

func TestParserRootClass(t *testing.T) {
	str := `Text {
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.ClassStatement
	if classStmt.Name != "Text" {
		t.Errorf("Incorrect class name, expected Text, take %q", classStmt.Name)
		return
	}
}

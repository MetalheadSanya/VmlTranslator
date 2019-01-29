package parser_test

import (
	"github.com/MetalheadSanya/VmlTranslator/parser"
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"strings"
	"testing"
)

func TestParserPropertyAssignment(t *testing.T) {
	str := `
Rectangle {
    width: 100
    height: 100
    color: "red"
	enabled: false
	clip: true	
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	root := stmt.ClassStatement
	if root.Name != "Rectangle" {
		t.Errorf("Incorrect root class parent, expected 'Rectangle', take %s", root.Name)
	}

	propertyAssignmentCount := len(root.PropertyAssignments)
	if propertyAssignmentCount != 5 {
		t.Errorf("Incorrect property assignments count, expected '5', take %d", propertyAssignmentCount)
	}

	if propertyAssignmentCount < 1 {
		return
	}
	prop1 := root.PropertyAssignments[0]
	if prop1.Property[0] != "width" {
		t.Errorf("Incorrect first property, expected 'width', take %s", prop1.Property[0])
	}
	arg1, ok := prop1.Expression.(*literals.IntegerLiteral)
	if ok && *arg1 != 100 {
		t.Errorf("Incorrect first property value, expected '100', take %d", *arg1)
	} else if !ok {
		t.Errorf("Incorrect first property type, expected IntegerLiteral")
	}

	if propertyAssignmentCount < 2 {
		return
	}
	prop2 := root.PropertyAssignments[1]
	if prop2.Property[0] != "height" {
		t.Errorf("Incorrect second property, expected 'height', take %s", prop2.Property[0])
	}
	arg2, ok := prop2.Expression.(*literals.IntegerLiteral)
	if ok && *arg2 != 100 {
		t.Errorf("Incorrect second property value, expected '100', take %d", *arg2)
	} else if !ok {
		t.Errorf("Incorrect second property type, expected IntegerLiteral")
	}

	if propertyAssignmentCount < 3 {
		return
	}
	prop3 := root.PropertyAssignments[2]
	if prop3.Property[0] != "color" {
		t.Errorf("Incorrect third property, expected 'color', take %s", prop3.Property[0])
	}
	arg3, ok := prop3.Expression.(*literals.StringLiteral)
	if ok && *arg3 != "red" {
		t.Errorf("Incorrect third property value, expected 'red', take %s", *arg3)
	} else if !ok {
		t.Errorf("Incorrect third property type, expected StringLiteral")
	}

	if propertyAssignmentCount < 4 {
		return
	}
	prop4 := root.PropertyAssignments[3]
	if prop4.Property[0] != "enabled" {
		t.Errorf("Incorrect fourth property, expected 'enabled', take %s", prop4.Property[0])
	}
	arg4, ok := prop4.Expression.(*literals.BooleanLiteral)
	if ok && *arg4 != false {
		t.Errorf("Incorrect fourth property value, expected 'false', take %t", *arg4)
	} else if !ok {
		t.Errorf("Incorrect fourth property type, expected BooleanLiteral")
	}

	if propertyAssignmentCount < 5 {
		return
	}
	prop5 := root.PropertyAssignments[4]
	if prop5.Property[0] != "clip" {
		t.Errorf("Incorrect fifth property, expected 'color', take %s", prop5.Property[0])
	}
	arg5, ok := prop5.Expression.(*literals.BooleanLiteral)
	if ok && *arg5 != false {
		t.Errorf("Incorrect fifth property value, expected 'red', take %t", *arg5)
	} else if !ok {
		t.Errorf("Incorrect fifth property type, expected BooleanLiteral")
	}
}

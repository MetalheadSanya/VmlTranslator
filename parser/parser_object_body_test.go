package parser

import (
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
	"strings"
	"testing"
)

func TestObjectBodyProperties(t *testing.T) {
	str := `{
a: 1
b: "3"
}`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseObjectBody()
	if err != nil {
		t.Error(err)
		return
	}

	checkObjectProperties(t, stmt)
}

func TestObjectBodyPropertiesSingleLine(t *testing.T) {
	str := `{ a: 1; b: "3" }`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseObjectBody()
	if err != nil {
		t.Error(err)
		return
	}

	checkObjectProperties(t, stmt)
}

func checkObjectProperties(t *testing.T, object *statement.Object) {
	if object.Name.Len() != 0 {
		t.Errorf("expected object name be empty")
	}

	if object.Children.Len() != 0 {
		t.Errorf("expected Children be empty")
	}

	if object.PropertyAssignments.Len() != 2 {
		t.Errorf("expected 2 property assignment")
	}

	iter := object.PropertyAssignments.Front()
	if iter == nil {
		return
	}

	prop1, ok := iter.Value.(*statement.PropertyAssignment)
	if !ok {
		t.Errorf("expected statement.PropertyAssignment")
	} else {
		if prop1.Property.Len() != 1 {
			t.Errorf("invalit property assigment path")
		} else if *prop1.Property.Front().Value.(*string) != "a" {
			t.Errorf("expected fisrt property name 'a', take %q", *prop1.Property.Front().Value.(*string))
		}

		expr, ok := prop1.Expression.(*literals.IntegerLiteral)
		if !ok {
			t.Errorf("expected fisrt property value type literals.IntegerLiteral")
		} else if *expr != 1 {
			t.Errorf("exprected first property value '1', take %d", *expr)
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}

	prop2, ok := iter.Value.(*statement.PropertyAssignment)
	if !ok {
		t.Errorf("expected statement.PropertyAssignment")
	} else {
		if prop2.Property.Len() != 1 {
			t.Errorf("invalit property assigment path")
		} else if *prop2.Property.Front().Value.(*string) != "b" {
			t.Errorf("expected second property name 'b', take %q", *prop2.Property.Front().Value.(*string))
		}

		expr, ok := prop2.Expression.(*literals.StringLiteral)
		if !ok {
			t.Errorf("expected fisrt property value type literals.StringLiteral")
		} else if *expr != "3" {
			t.Errorf("exprected first property value '3', take %q", *expr)
		}
	}
}

func TestObjectBodyChild(t *testing.T) {
	str := `{
Rect { }
Item { }
}`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseObjectBody()
	if err != nil {
		t.Error(err)
		return
	}

	if stmt.Name.Len() != 0 {
		t.Errorf("expected object name be empty")
	}

	if stmt.Children.Len() != 2 {
		t.Errorf("expected 2 Children")
	}

	if stmt.PropertyAssignments.Len() != 0 {
		t.Errorf("expected empty property assignment")
	}

	iter := stmt.Children.Front()
	if iter == nil {
		return
	}

	child1, ok := iter.Value.(*statement.Object)
	if !ok {
		t.Errorf("expected statement.Object")
	} else {
		if child1.Name.Len() != 1 {
			t.Errorf("expected first child object name be Rect, take Empty")
		} else if *child1.Name.Front().Value.(*string) != "Rect" {
			t.Errorf("expected first child object name Rect, take %q", *child1.Name.Front().Value.(*string))
		}
		if child1.Children.Len() != 0 {
			t.Errorf("exprected empty Children")
		}
		if child1.PropertyAssignments.Len() != 0 {
			t.Errorf("expected emptyPropertyAssignments")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}

	child2, ok := iter.Value.(*statement.Object)
	if !ok {
		t.Errorf("expected statement.Object")
	} else {
		if child2.Name.Len() != 1 {
			t.Errorf("expected second child object name be Rect, take Empty")
		} else if *child2.Name.Front().Value.(*string) != "Item" {
			t.Errorf("expected second child object name Rect, take %q", *child2.Name.Front().Value.(*string))
		}
		if child2.Children.Len() != 0 {
			t.Errorf("exprected empty Children")
		}
		if child2.PropertyAssignments.Len() != 0 {
			t.Errorf("expected emptyPropertyAssignments ")
		}
	}
}

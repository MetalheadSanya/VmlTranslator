package parser

import (
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
	"strings"
	"testing"
)

func TestIntListLiteral(t *testing.T) {
	str := `[1, 2, 3]`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseListLiteralExpression()
	if err != nil {
		t.Error(err)
		return
	}

	elementsCount := stmt.Len()
	if elementsCount != 3 {
		t.Errorf("Incorrect elements count, expected '3', take %d", elementsCount)
	}

	element := stmt.Front()
	if element == nil {
		return
	}
	el1, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el1 != 1 {
		t.Errorf("Incorrect first property value, expected '1', take %d", *el1)
	} else if !ok {
		t.Errorf("Incorrect first property type, expected IntegerLiteral")
	}

	element = element.Next()
	if element == nil {
		return
	}
	el2, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el2 != 2 {
		t.Errorf("Incorrect second property value, expected '2', take %d", *el2)
	} else if !ok {
		t.Errorf("Incorrect second property type, expected IntegerLiteral")
	}

	element = element.Next()
	if element == nil {
		return
	}
	el3, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el3 != 3 {
		t.Errorf("Incorrect third property value, expected '3', take %d", *el3)
	} else if !ok {
		t.Errorf("Incorrect third property type, expected IntegerLiteral")
	}
}

func TestEmptyListLiteral(t *testing.T) {
	str := `[]`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseListLiteralExpression()
	if err != nil {
		t.Error(err)
		return
	}

	elementsCount := stmt.Len()
	if elementsCount != 0 {
		t.Errorf("Incorrect elements count, expected '0', take %d", elementsCount)
	}
}

func TestMultilineListLiteral(t *testing.T) {
	str := `[1, 

2,
3
,
]`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseListLiteralExpression()
	if err != nil {
		t.Error(err)
		return
	}

	elementsCount := stmt.Len()
	if elementsCount != 3 {
		t.Errorf("Incorrect elements count, expected '3', take %d", elementsCount)
	}

	element := stmt.Front()
	if element == nil {
		return
	}
	el1, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el1 != 1 {
		t.Errorf("Incorrect first property value, expected '1', take %d", *el1)
	} else if !ok {
		t.Errorf("Incorrect first property type, expected IntegerLiteral")
	}

	element = element.Next()
	if element == nil {
		return
	}
	el2, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el2 != 2 {
		t.Errorf("Incorrect second property value, expected '2', take %d", *el2)
	} else if !ok {
		t.Errorf("Incorrect second property type, expected IntegerLiteral")
	}

	element = element.Next()
	if element == nil {
		return
	}
	el3, ok := element.Value.(*literals.IntegerLiteral)
	if ok && *el3 != 3 {
		t.Errorf("Incorrect third property value, expected '3', take %d", *el3)
	} else if !ok {
		t.Errorf("Incorrect third property type, expected IntegerLiteral")
	}
}

func TestExpressionListLiteral(t *testing.T) {
	str := `[foo.baz(bar)]`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseListLiteralExpression()
	if err != nil {
		t.Error(err)
		return
	}

	elementsCount := stmt.Len()
	if elementsCount != 1 {
		t.Errorf("Incorrect elements count, expected '1', take %d", elementsCount)
	}

	element := stmt.Front()
	if element == nil {
		return
	}
	functionCallExpression, ok := element.Value.(*statement.FunctionCallExpression)
	if !ok {
		t.Errorf("Incorrect expression type, expected FunctionCallExpression")
		return
	}
	memberExpression, ok := functionCallExpression.Expression.(*statement.ExplicitMemberExpression)
	if !ok {
		t.Errorf("Incorrect expression type, expected ExplicitMemberExpression")
		return
	}
	// method name
	if memberExpression.Member != "baz" {
		t.Errorf("Incorrect member value, expected 'baz', take %s", memberExpression.Member)
	}

	// variable name
	ident, ok := memberExpression.Expression.(*statement.Identifier)
	if ok && *ident != "foo" {
		t.Errorf("Incorrect left identifier, expected 'foo', take %s", *ident)
	} else if !ok {
		t.Errorf("Incorrect expression type, expected Identifier")
	}

	// args

	argCount := len(functionCallExpression.ArgumentList)
	if argCount != 1 {
		t.Errorf("Incorrect arguments count, expected '1', take %d", argCount)
		return
	}
	arg1 := functionCallExpression.ArgumentList[0]
	ident1, ok := arg1.(*statement.Identifier)
	if ok && *ident1 != "bar" {
		t.Errorf("Incorrect argument value, expected 'bar', take %s", *ident1)
	} else if !ok {
		t.Errorf("Incorrect argument type, expected Identifier")
	}
}

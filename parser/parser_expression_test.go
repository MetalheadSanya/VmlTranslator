package parser

import (
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
	"strings"
	"testing"
)

func TestParserFunctionCallWithLiterals(t *testing.T) {
	str := `(12, "test", 13.0)`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseFunctionCallExpression()
	if err != nil {
		t.Error(err)
		return
	}
	argCount := len(stmt.ArgumentList)
	if argCount != 3 {
		t.Errorf("Incorrect arguments count, expected '3', take %d", argCount)
	}
	if argCount < 1 {
		return
	}
	arg1 := stmt.ArgumentList[0]
	intArg, ok := arg1.(*literals.IntegerLiteral)
	if ok && *intArg != 12 {
		t.Errorf("Incorrect first argument value, expected '12', take %d", *intArg)
	} else if !ok {
		t.Errorf("Incorrect first argument type, expected IntegerLiteral")
	}
	if argCount < 2 {
		return
	}
	arg2 := stmt.ArgumentList[1]
	strArg, ok := arg2.(*literals.StringLiteral)
	if ok && *strArg != "test" {
		t.Errorf("Incorrect second argument value, expected 'test', take %s", *strArg)
	} else if !ok {
		t.Errorf("Incorrect second argument type, expected StringLiteral")
	}
	if argCount < 3 {
		return
	}
	arg3 := stmt.ArgumentList[2]
	fpArg, ok := arg3.(*literals.FloatingPointerLiteral)
	if ok && *fpArg != 13.0 {
		t.Errorf("Incorrect second argument value, expected 'test', take %f", *fpArg)
	} else if !ok {
		t.Errorf("Incorrect second argument type, expected FloatingPointerLiteral")
	}
}

func TestParserFunctionCallWithIdentifier(t *testing.T) {
	str := `(foo, baz)`
	r := strings.NewReader(str)
	p := VmlParser(r)
	stmt, err := p.parseFunctionCallExpression()
	if err != nil {
		t.Error(err)
		return
	}
	argCount := len(stmt.ArgumentList)
	if argCount != 2 {
		t.Errorf("Incorrect arguments count, expected '2', take %d", argCount)
	}
	if argCount < 1 {
		return
	}
	arg1 := stmt.ArgumentList[0]
	ident1, ok := arg1.(*statement.Identifier)
	if ok && *ident1 != "foo" {
		t.Errorf("Incorrect first argument value, expected 'foo', take %s", *ident1)
	} else if !ok {
		t.Errorf("Incorrect first argument type, expected Identifier")
	}
	if argCount < 2 {
		return
	}
	arg2 := stmt.ArgumentList[1]
	ident2, ok := arg2.(*statement.Identifier)
	if ok && *ident2 != "baz" {
		t.Errorf("Incorrect second argument value, expected 'baz', take %s", *ident2)
	} else if !ok {
		t.Errorf("Incorrect second argument type, expected Identifier")
	}
}

func TestParserIdentifier(t *testing.T) {
	str := `foo`
	r := strings.NewReader(str)
	p := VmlParser(r)

	ident, err := p.parseIdentifier()
	if err != nil {
		t.Error(err)
		return
	}
	if *ident != "foo" {
		t.Errorf("Incorrect identifier, expected 'foo', take %s", *ident)
	}

	r = strings.NewReader(str)
	p = VmlParser(r)
	stmt, err := p.parseExpression()
	if err != nil {
		t.Error(err)
		return
	}
	ident, ok := stmt.(*statement.Identifier)
	if ok && *ident != "foo" {
		t.Errorf("Incorrect identifier, expected 'foo', take %s", *ident)
	} else if !ok {
		t.Errorf("Incorrect expression type, expected Identifier")
	}
}

func TestParserProperty(t *testing.T) {
	str := `.foo`
	r := strings.NewReader(str)
	p := VmlParser(r)

	memberExpression, err := p.parseExplicitMemberExpression()
	if err != nil {
		t.Error(err)
		return
	}
	if memberExpression.Member != "foo" {
		t.Errorf("Incorrect member value, expected 'foo', take %s", memberExpression.Member)
	}
}

func TestParserPropertyFull(t *testing.T) {
	str := `foo.baz`
	r := strings.NewReader(str)
	p := VmlParser(r)

	expr, err := p.parseExpression()
	if err != nil {
		t.Error(err)
		return
	}
	memberExpression, ok := expr.(*statement.ExplicitMemberExpression)
	if !ok {
		t.Errorf("Incorrect expression type, expected ExplicitMemberExpression")
		return
	}
	if memberExpression.Member != "baz" {
		t.Errorf("Incorrect member value, expected 'baz', take %s", memberExpression.Member)
	}
	ident, ok := memberExpression.Expression.(*statement.Identifier)
	if ok && *ident != "foo" {
		t.Errorf("Incorrect left identifier, expected 'foo', take %s", *ident)
	} else if !ok {
		t.Errorf("Incorrect expression type, expected Identifier")
	}
}

func TestParserFunctionCallFull(t *testing.T) {
	str := `foo.baz(bar)`
	r := strings.NewReader(str)
	p := VmlParser(r)

	expr, err := p.parseExpression()
	if err != nil {
		t.Error(err)
		return
	}

	functionCallExpression, ok := expr.(*statement.FunctionCallExpression)
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

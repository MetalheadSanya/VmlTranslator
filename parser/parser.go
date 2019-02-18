package parser

import (
	"container/list"
	"fmt"
	"github.com/MetalheadSanya/VmlTranslator/lexer"
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
	"io"
	"strconv"
)

type Parser struct {
	s   *lexer.Scanner
	buf struct {
		tok lexer.Token
		lit string
		n   int
	}
}

func VmlParser(r io.Reader) *Parser {
	return &Parser{s: lexer.VmlScanner(r)}
}

func (p *Parser) scan() (tok lexer.Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreNewLine() (tok lexer.Token, lit string) {
	tok, lit = p.scan()
	if tok == lexer.NewLine {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) Parse() (*statement.File, error) {
	stmt := &statement.File{}

	for {
		tok, _ := p.scanIgnoreNewLine()

		if tok == lexer.Import {
			p.unscan()
			importStmt, err := p.parseImportStatement()
			if err != nil {
				return nil, err
			}
			stmt.Imports = append(stmt.Imports, importStmt)
		} else if tok == lexer.Identifier && stmt.Class.Name == "" {
			p.unscan()
			classStmt, err := p.parseRootClassStatement()
			if err != nil {
				return nil, err
			}
			stmt.Class = *classStmt
		} else {
			p.unscan()
			break
		}
	}

	return stmt, nil
}

func (p *Parser) parseImportStatement() (interface{}, error) {
	var stmt interface{}

	tok, lit := p.scan()
	if tok != lexer.Import {
		return nil, fmt.Errorf("found %q, expected import keyword", lit)
	}
	tok, lit = p.scan()
	if tok == lexer.Identifier {
		namespace := &statement.ImportNamespace{}
		for {
			if tok != lexer.Identifier {
				return nil, fmt.Errorf("found %q, expected identifier", lit)
			}
			namespace.ModuleIdentifier = append(namespace.ModuleIdentifier, lit)

			tok, lit = p.scan()
			if tok != lexer.Dot {
				break
			}
			tok, lit = p.scan()
		}
		if tok != lexer.DoubleLiteral {
			return nil, fmt.Errorf("found %q, expected double literal", lit)
		}
		version, err := strconv.ParseFloat(lit, 32)
		if err != nil {
			return nil, err
		}
		namespace.Version = float32(version)
		stmt = namespace
	} else if tok == lexer.StringLiteral {
		directory := &statement.ImportDirectory{}
		directory.Directory = lit
		stmt = directory
	} else {
		return nil, fmt.Errorf("found %q, expected identifier or string literal", lit)
	}

	tok, lit = p.scan()
	if tok != lexer.As {
		p.unscan()
		return stmt, p.scanNewLineOrSkipEof()
	}

	tok, lit = p.scan()
	if tok != lexer.Identifier {
		return nil, fmt.Errorf("found %q, expected identifier", lit)
	}

	qualifier := lit

	switch stmt := stmt.(type) {
	case *statement.ImportDirectory:
		stmt.Qualifier = &qualifier
	case *statement.ImportNamespace:
		stmt.Qualifier = &qualifier
	}

	return stmt, p.scanNewLineOrSkipEof()
}

func (p *Parser) parseRootClassStatement() (*statement.Class, error) {
	stmt := statement.NewClass()
	tok, lex := p.scan()

	if tok != lexer.Identifier {
		return nil, fmt.Errorf("found %q, expexted identifier", lex)
	}
	stmt.Name = lex

	tok, lex = p.scan()
	if tok != lexer.LeftCurlyBracket {
		return nil, fmt.Errorf("found %q, expexted '{'", lex)
	}

	for {
		tok, lex = p.scanIgnoreNewLine()
		if tok == lexer.Identifier {
			p.unscan()
			path, err := p.parsePath()
			if err != nil {
				return nil, err
			}
			tok, lex = p.scan()
			if tok == lexer.Colon {
				expr, err := p.parseExpression()

				if err != nil {
					return nil, err
				}
				prop := &statement.PropertyAssignment{}
				prop.Property = path
				prop.Expression = expr

				stmt.PropertyAssignments = append(stmt.PropertyAssignments, *prop)
			} else if tok == lexer.LeftCurlyBracket {
				p.unscan()
				object, err := p.parseObjectBody()

				if err != nil {
					return nil, err
				}
				object.Name = path

				stmt.Children.PushBack(object)
			} else {
				return nil, fmt.Errorf("expected ':' or '{', take %q", lex)
			}
		} else if tok == lexer.Property || tok == lexer.Default {
			p.unscan()
			prop, err := p.parseDefinePropertyStatement()
			if err != nil {
				return nil, err
			}
			stmt.Properties.PushBack(prop)
		} else {
			break
		}
	}

	if tok != lexer.RightCurlyBracket {
		return nil, fmt.Errorf("expected '}', statement.Identifier, 'default' or 'property'"+
			", found %q", lex)
	}

	return stmt, nil
}

func (p *Parser) parsePropertyAssignmentStatement() (*statement.PropertyAssignment, error) {
	stmt := &statement.PropertyAssignment{}

	var tok lexer.Token
	var lex string

	path, err := p.parsePath()

	if err != nil {
		return nil, err
	}
	stmt.Property = path

	tok, lex = p.scan()
	if tok != lexer.Colon {
		return nil, fmt.Errorf("found %q, expected ':' or '.'", lex)
	}

	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	stmt.Expression = expr

	return stmt, nil
}

func (p *Parser) parseDefinePropertyStatement() (*statement.Property, error) {
	stmt := &statement.Property{}

	var tok lexer.Token
	var lex string

	tok, lex = p.scan()
	if tok == lexer.Default {
		stmt.IsDefault = true
		tok, lex = p.scan()
	} else {
		stmt.IsDefault = false
	}

	if tok != lexer.Property {
		return nil, fmt.Errorf("found %q, expeced property keyword", lex)
	}

	t, err := p.parseTypeStatement()
	if err != nil {
		return nil, err
	}

	stmt.PropertyType = t

	tok, lex = p.scan()
	if tok != lexer.Identifier {
		return nil, fmt.Errorf("found %q, expeced identifier", lex)
	}

	stmt.PropertyName = lex

	tok, lex = p.scan()
	if tok != lexer.Colon {
		p.unscan()
		return stmt, nil
	}

	expr, err := p.parseExpression()

	if err != nil {
		return nil, err
	}

	stmt.PropertyValue = expr

	err = p.scanNewLineOrSkipEof()

	return stmt, err
}

func (p *Parser) parseExpression() (interface{}, error) {
	var stmt interface{}
	var err error

	for {
		tok, _ := p.scan()
		p.unscan()
		if tok == lexer.NewLine ||
			tok == lexer.Eof ||
			tok == lexer.Comma ||
			tok == lexer.RightParenthesis ||
			tok == lexer.RightSquareBracket ||
			tok == lexer.RightCurlyBracket {
			break
		}

		if tok == lexer.Semicolon {
			p.scan()
			break
		}

		if tok == lexer.IntLiteral ||
			tok == lexer.StringLiteral ||
			tok == lexer.DoubleLiteral ||
			tok == lexer.True ||
			tok == lexer.False {
			var literal interface{}
			literal, err = p.parseLiteralExpression()
			if err == nil {
				stmt = literal
				continue
			}
		}

		if tok == lexer.LeftSquareBracket {
			var literal interface{}
			literal, err = p.parseListLiteralExpression()
			if err == nil {
				stmt = literal
				continue
			}
		}

		if tok == lexer.Identifier {
			var ident *statement.Identifier
			ident, err = p.parseIdentifier()
			if err == nil {
				stmt = ident
				continue
			}
		}

		if tok == lexer.Dot && stmt != nil {
			var explicitMember *statement.ExplicitMemberExpression
			explicitMember, err = p.parseExplicitMemberExpression()
			// check on custom error
			if err == nil {
				explicitMember.Expression = stmt
				stmt = explicitMember
				continue
			}
		}

		if tok == lexer.LeftParenthesis && stmt != nil {
			var functionCall *statement.FunctionCallExpression
			functionCall, err = p.parseFunctionCallExpression()
			// check on custom error
			if err == nil {
				functionCall.Expression = stmt
				stmt = functionCall
				continue
			}
		}
	}

	return stmt, nil
}

func (p *Parser) parseLiteralExpression() (interface{}, error) {
	var stmt interface{}

	var tok lexer.Token
	var lex string

	tok, lex = p.scan()
	switch tok {
	case lexer.StringLiteral:
		str := literals.StringLiteral(lex)
		stmt = &str
	case lexer.IntLiteral:
		str := lex
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		lit := literals.IntegerLiteral(val)
		stmt = &lit
	case lexer.DoubleLiteral:
		str := lex
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		lit := literals.FloatingPointerLiteral(val)
		stmt = &lit
	case lexer.True:
		lit := literals.BooleanLiteral(true)
		stmt = &lit
	case lexer.False:
		lit := literals.BooleanLiteral(false)
		stmt = &lit
	}

	if stmt == nil {
		p.unscan()
		return nil, fmt.Errorf("not found literal expression")
	}

	return stmt, nil
}

func (p *Parser) parseListLiteralExpression() (literals.ListLiteral, error) {
	stmt := list.New()

	tok, lex := p.scan()
	if tok != lexer.LeftSquareBracket {
		return nil, fmt.Errorf("expected '[' take %q", lex)
	}

	for {
		tok, lex = p.scanIgnoreNewLine()
		if tok == lexer.RightSquareBracket {
			return stmt, nil
		}
		p.unscan()
		element, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.PushBack(element)
		tok, lex = p.scanIgnoreNewLine()
		if tok == lexer.Comma {
			p.scanIgnoreNewLine()
			p.unscan()
			continue
		}
		if tok == lexer.RightSquareBracket {
			break
		}

		return nil, fmt.Errorf("expected ']' or ',', take %q", lex)
	}

	return stmt, nil
}

func (p *Parser) parseIdentifier() (*statement.Identifier, error) {
	tok, lex := p.scan()
	if tok == lexer.Identifier {
		return &lex, nil
	}
	p.unscan()
	return nil, fmt.Errorf("not found identifier")
}

func (p *Parser) parseExplicitMemberExpression() (*statement.ExplicitMemberExpression, error) {
	stmt := &statement.ExplicitMemberExpression{}

	var tok lexer.Token
	var lex string

	tok, lex = p.scan()
	if tok != lexer.Dot {
		p.unscan()
		return nil, fmt.Errorf("found %q, exprexted '.'", lex)
	}

	tok, lex = p.scan()
	if tok != lexer.Identifier {
		// TODO: custom error
		p.unscan()
		return nil, fmt.Errorf("found %q, exprexted identifier", lex)
	}

	stmt.Member = statement.Identifier(lex)
	return stmt, nil
}

func (p *Parser) parseFunctionCallExpression() (*statement.FunctionCallExpression, error) {
	stmt := &statement.FunctionCallExpression{}

	var tok lexer.Token
	var lex string

	tok, lex = p.scan()
	if tok != lexer.LeftParenthesis {
		p.unscan()
		return nil, fmt.Errorf("found %q, exprexted '('", lex)
	}

	tok, lex = p.scan()
	if tok == lexer.RightParenthesis {
		return stmt, nil
	}
	p.unscan()

	for {
		arg, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		stmt.ArgumentList = append(stmt.ArgumentList, arg)

		tok, lex = p.scan()
		if tok == lexer.Comma {
			continue
		}
		if tok == lexer.RightParenthesis {
			break
		}
		return nil, fmt.Errorf("expected ',' or ')' or expression")
	}

	return stmt, nil
}

func (p *Parser) parseTypeStatement() (interface{}, error) {
	tok, lex := p.scan()

	if tok != lexer.Identifier &&
		tok != lexer.StringType &&
		tok != lexer.VarType &&
		tok != lexer.DoubleType &&
		tok != lexer.BoolType &&
		tok != lexer.ListType {
		return nil, fmt.Errorf("expected type, take %q", lex)
	}

	name := lex

	if tok == lexer.Identifier || tok == lexer.ListType {
		basicType := tok
		tok, lex = p.scan()

		if tok == lexer.LessThanSign {
			stmt := &statement.GenericType{}
			childType, err := p.parseTypeStatement()
			if err != nil {
				return nil, err
			}
			stmt.Container = name
			stmt.Element = childType
			tok, lex = p.scan()
			if tok != lexer.MoreThanSign {
				return nil, fmt.Errorf("expected '>', take %q", lex)
			}
			return stmt, nil
		} else if basicType == lexer.ListType {
			return nil, fmt.Errorf("expected '<', take %q", lex)
		} else {
			p.unscan()
			return &name, nil
		}
	} else {
		return &name, nil
	}
}

func (p *Parser) parseObjectBody() (*statement.Object, error) {
	stmt := statement.NewObject()

	tok, lex := p.scan()

	if tok != lexer.LeftCurlyBracket {
		return nil, errorf(lexer.LeftCurlyBracket, lex)
	}

	for {
		tok, lex = p.scanIgnoreNewLine()

		if tok == lexer.Identifier {
			p.unscan()
		} else if tok == lexer.RightCurlyBracket {
			break
		}

		path, err := p.parsePath()
		if err != nil {
			return nil, err
		}

		tok, lex = p.scan()
		if tok == lexer.Colon {
			expr, err := p.parseExpression()

			if err != nil {
				return nil, err
			}
			prop := &statement.PropertyAssignment{}
			prop.Property = path
			prop.Expression = expr

			stmt.PropertyAssignments.PushBack(prop)
		} else if tok == lexer.LeftCurlyBracket {
			p.unscan()
			child, err := p.parseObjectBody()

			if err != nil {
				return nil, err
			}

			child.Name = path
			stmt.Children.PushBack(child)
		} else {
			return nil, errorf(lexer.Colon, lex)
		}
	}

	return stmt, nil
}

func errorf(expectedToken lexer.Token, lex string) error {
	return fmt.Errorf("expected %#v, take %q", expectedToken, lex)
}

func (p *Parser) parsePath() (*list.List, error) {
	var tok lexer.Token
	var lex string

	stmt := list.New()
	for {
		tok, lex = p.scan()
		if tok != lexer.Identifier {
			return nil, errorf(lexer.Identifier, lex)
		}
		ident := lex
		stmt.PushBack(&ident)

		tok, lex = p.scan()
		if tok != lexer.Dot {
			p.unscan()
			break
		}
	}

	return stmt, nil
}

func (p *Parser) scanNewLineOrSkipEof() error {
	tok, lit := p.scan()

	if tok != lexer.NewLine && tok != lexer.Eof {
		return fmt.Errorf("found %q, expected new line or end of file", lit)
	}
	if tok == lexer.Eof {
		p.unscan()
	}

	return nil
}

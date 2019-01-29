package parser

import (
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
			stmt.NamespaceImports = append(stmt.NamespaceImports, importStmt)
		} else if tok == lexer.Identifier && stmt.ClassStatement.Name == "" {
			p.unscan()
			classStmt, err := p.parseRootClassStatement()
			if err != nil {
				return nil, err
			}
			stmt.ClassStatement = *classStmt
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

func (p *Parser) parseRootClassStatement() (*statement.RootClass, error) {
	stmt := &statement.RootClass{}
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
		if tok != lexer.Identifier {
			break
		}
		p.unscan()
		expr, err := p.parsePropertyAssignmentStatement()
		if err == nil {
			stmt.PropertyAssignments = append(stmt.PropertyAssignments, *expr)
			continue
		}
	}

	if tok != lexer.RightCurlyBracket {
		return nil, fmt.Errorf("found %q, expexted '}'", lex)
	}

	return stmt, nil
}

func (p *Parser) parsePropertyAssignmentStatement() (*statement.PropertyAssignment, error) {
	stmt := &statement.PropertyAssignment{}

	var tok lexer.Token
	var lex string

	for {
		tok, lex = p.scan()
		if tok != lexer.Identifier {
			return nil, fmt.Errorf("found %q, expexted '.'", lex)
		}
		stmt.Property = append(stmt.Property, lex)

		tok, lex = p.scan()
		if tok != lexer.Dot {
			p.unscan()
			break
		}
	}

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

func (p *Parser) parseExpression() (interface{}, error) {
	var stmt interface{}
	var err error

	for {
		tok, _ := p.scan()
		p.unscan()
		if tok == lexer.NewLine || tok == lexer.Eof || tok == lexer.Comma || tok == lexer.RightParenthesis {
			break
		}

		if tok == lexer.IntLiteral || tok == lexer.StringLiteral || tok == lexer.DoubleLiteral {
			var literal interface{}
			literal, err = p.parseLiteralExpression()
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

func (p *Parser) parseIdentifier() (*statement.Identifier, error) {
	tok, lex := p.scan()
	if tok == lexer.Identifier {
		ident := statement.Identifier(lex)
		return &ident, nil
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

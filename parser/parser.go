package parser

import (
	"fmt"
	"github.com/MetalheadSanya/VmlTranslator/lexer"
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

func (p *Parser) Parse() (*FileStatement, error) {
	stmt := &FileStatement{}

	for {
		tok, _ := p.scanIgnoreNewLine()

		if tok == lexer.Import {
			p.unscan()
			importStmt, err := p.parseImportStatement()
			if err != nil {
				return nil, err
			}
			stmt.NamespaceImports = append(stmt.NamespaceImports, *importStmt)
		} else {
			p.unscan()
			break
		}
	}

	return stmt, nil
}

func (p *Parser) parseImportStatement() (statement *ImportNamespaceStatement, err error) {
	stmt := &ImportNamespaceStatement{}

	tok, lit := p.scan()
	if tok != lexer.Import {
		return nil, fmt.Errorf("found %q, expected import keyword", lit)
	}
	for {
		tok, lit = p.scan()
		if tok != lexer.Identifier {
			return nil, fmt.Errorf("found %q, expected identifier", lit)
		}
		stmt.ModuleIdentifier = append(stmt.ModuleIdentifier, lit)

		tok, lit = p.scan()
		if tok != lexer.Dot {
			break
		}
	}
	if tok != lexer.DoubleLiteral {
		return nil, fmt.Errorf("found %q, expected double literal", lit)
	}
	version, err := strconv.ParseFloat(lit, 32)
	if err != nil {
		return nil, err
	}
	stmt.Version = float32(version)
	tok, lit = p.scan()

	if tok != lexer.As {
		p.unscan()
		return stmt, p.scanNewLineOrSkipEof()
	}

	tok, lit = p.scan()
	if tok != lexer.Identifier {
		return nil, fmt.Errorf("found %q, expected identifier", lit)
	}

	stmt.Qualifier = &lit

	return stmt, p.scanNewLineOrSkipEof()
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

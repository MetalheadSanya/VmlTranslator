package statement

import "github.com/MetalheadSanya/VmlTranslator/parser/literals"

// Is any literal
type LiteralExpression interface{}

func IsLiteralExpression(e interface{}) bool {
	return literals.IsLiteral(e)
}

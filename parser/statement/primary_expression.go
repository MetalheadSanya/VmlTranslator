package statement

// Is any literal-expression
type PrimaryExpression interface{}

func IsPrimaryExpression(e interface{}) bool {
	return IsLiteralExpression(e)
}

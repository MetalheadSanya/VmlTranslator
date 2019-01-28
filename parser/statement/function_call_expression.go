package statement

// Must be any expression or identifier
type FunctionCallArgumentList []interface{}

type FunctionCallExpression struct {
	Expression   interface{}
	ArgumentList FunctionCallArgumentList
}

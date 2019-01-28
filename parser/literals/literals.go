package literals

type BooleanLiteral bool
type FloatingPointerLiteral float64
type IntegerLiteral int
type StringLiteral string

func IsLiteral(l interface{}) bool {
	switch l.(type) {
	case BooleanLiteral, FloatingPointerLiteral, IntegerLiteral, StringLiteral:
		return true
	default:
		return false
	}
}

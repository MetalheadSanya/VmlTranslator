package literals

import "container/list"

type BooleanLiteral bool
type FloatingPointerLiteral float64
type IntegerLiteral int
type StringLiteral string
type ListLiteral = *list.List

func IsLiteral(l interface{}) bool {
	switch l.(type) {
	case BooleanLiteral, FloatingPointerLiteral, IntegerLiteral, StringLiteral, ListLiteral:
		return true
	default:
		return false
	}
}

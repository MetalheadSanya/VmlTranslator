package statement

import "container/list"

type PropertyAssignment struct {
	Property   *list.List
	Expression interface{}
}

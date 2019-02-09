package statement

import "container/list"

type Object struct {
	Name                *list.List
	PropertyAssignments *list.List
	Children            *list.List
}

func NewObject() *Object {
	return &Object{list.New(), list.New(), list.New()}
}

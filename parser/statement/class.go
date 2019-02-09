package statement

import "container/list"

type Class struct {
	Name                string
	PropertyAssignments []PropertyAssignment
	Children            *list.List
	Properties          *list.List
}

func NewClass() *Class {
	return &Class{"", nil, list.New(), list.New()}
}

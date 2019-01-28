package statement

type PropertyAssignment struct {
	Property   PropertyPath
	Expression interface{}
}

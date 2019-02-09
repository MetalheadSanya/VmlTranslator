package statement

type Property struct {
	IsDefault     bool
	PropertyType  interface{}
	PropertyName  string
	PropertyValue interface{}
}

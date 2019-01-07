package parser

type ImportStatement interface {
	Qualifier() *string

	setQualifier(q *string)
}

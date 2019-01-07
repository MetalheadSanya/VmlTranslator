package parser

type ImportDirectoryStatement struct {
	directory string
	qualifier *string
}

func (s ImportDirectoryStatement) Directory() string {
	return s.directory
}

func (s ImportDirectoryStatement) Qualifier() *string {
	return s.qualifier
}

func (s *ImportDirectoryStatement) setQualifier(q *string) {
	s.qualifier = q
}

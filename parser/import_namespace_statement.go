package parser

type ImportNamespaceStatement struct {
	moduleIdentifier []string
	version          float32
	qualifier        *string
}

func (s ImportNamespaceStatement) ModuleIdentifier() []string {
	return s.moduleIdentifier
}

func (s ImportNamespaceStatement) Version() float32 {
	return s.version
}

func (s *ImportNamespaceStatement) setQualifier(q *string) {
	s.qualifier = q
}

func (s ImportNamespaceStatement) Qualifier() *string {
	return s.qualifier
}

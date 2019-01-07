package parser

type FileStatement struct {
	namespaceImports []ImportStatement
}

func (f FileStatement) NamespaceImports() []ImportStatement {
	return f.namespaceImports
}

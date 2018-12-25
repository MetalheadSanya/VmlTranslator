package parser

type ImportNamespaceStatement struct {
	ModuleIdentifier []string
	Version          float32
	Qualifier        *string
}

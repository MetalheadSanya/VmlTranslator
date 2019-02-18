package walker

import (
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
)

type Listener interface {
	EnterImportNamespace(namespace *statement.ImportNamespace)
	ExitImportNamespace(namespace *statement.ImportNamespace)
	EnterImportDirectory(directory *statement.ImportDirectory)
	ExitImportDirectory(directory *statement.ImportDirectory)
	EnterClass(class *statement.Class)
	ExitClass(class *statement.Class)
	EnterPropertyAssignment(assignment *statement.PropertyAssignment)
	ExitPropertyAssignment(assignment *statement.PropertyAssignment)
	EnterProperty(property *statement.Property)
	ExitProperty(property *statement.Property)
	EnterChild(object *statement.Object)
	ExitChild(object *statement.Object)
	EnterExplicitMember(expression *statement.ExplicitMemberExpression)
	ExitExplicitMember(expression *statement.ExplicitMemberExpression)
	EnterFunctionCall(expression *statement.FunctionCallExpression)
	ExitFunctionCall(expression *statement.FunctionCallExpression)
	EnterArgumentList(list statement.FunctionCallArgumentList)
	ExitArgumentList(list statement.FunctionCallArgumentList)
	EnterInteger(literal *literals.IntegerLiteral)
	ExitInteger(literal *literals.IntegerLiteral)
	EnterFloatPointer(literal *literals.FloatingPointerLiteral)
	ExitFloatPointer(literal *literals.FloatingPointerLiteral)
	EnterString(literal *literals.StringLiteral)
	ExitString(literal *literals.StringLiteral)
	EnterBoolean(literal *literals.BooleanLiteral)
	ExitBoolean(literal *literals.BooleanLiteral)
	EnterList(literal literals.ListLiteral)
	ExitList(literal literals.ListLiteral)
}

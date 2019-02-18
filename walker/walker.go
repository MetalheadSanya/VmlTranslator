package walker

import (
	"fmt"
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
)

type Walker struct {
	file     *statement.File
	listener Listener
}

func NewWalker(file *statement.File, listener Listener) Walker {
	return Walker{file, listener}
}

func (w Walker) Walk() error {
	for _, stmt := range w.file.Imports {
		switch i := stmt.(type) {
		case *statement.ImportDirectory:
			w.listener.EnterImportDirectory(i)
			w.listener.ExitImportDirectory(i)
		case *statement.ImportNamespace:
			w.listener.EnterImportNamespace(i)
			w.listener.ExitImportNamespace(i)
		default:
			return fmt.Errorf("unsupported import statement")
		}
	}

	if w.file != nil {
		err := w.visitClass(&w.file.Class)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w Walker) visitClass(class *statement.Class) error {
	w.listener.EnterClass(class)
	iter := class.Properties.Front()
	for iter != nil {
		stmt, ok := iter.Value.(*statement.Property)
		if !ok {
			return fmt.Errorf("unsupported property statement")
		}
		err := w.visitProperty(stmt)
		if err != nil {
			return err
		}
		iter = iter.Next()
	}
	for _, stmt := range class.PropertyAssignments {
		err := w.visitPropertyAssignment(&stmt)
		if err != nil {
			return err
		}
	}
	iter = class.Children.Front()
	for iter != nil {
		stmt, ok := iter.Value.(*statement.Object)
		if !ok {
			return fmt.Errorf("unsupported child statement")
		}
		err := w.visitObject(stmt)
		if err != nil {
			return err
		}
		iter = iter.Next()
	}
	w.listener.ExitClass(class)
	return nil
}

func (w Walker) visitObject(object *statement.Object) error {
	w.listener.EnterChild(object)
	iter := object.Children.Front()
	for iter != nil {
		stmt, ok := iter.Value.(*statement.Object)
		if !ok {
			return fmt.Errorf("unsupported child statement")
		}
		err := w.visitObject(stmt)
		if err != nil {
			return err
		}
		iter = iter.Next()
	}
	iter = object.PropertyAssignments.Front()
	for iter != nil {
		stmt, ok := iter.Value.(*statement.PropertyAssignment)
		if !ok {
			return fmt.Errorf("unsupported property assignment statement")
		}
		err := w.visitPropertyAssignment(stmt)
		if err != nil {
			return err
		}
		iter = iter.Next()
	}
	w.listener.ExitChild(object)
	return nil
}

func (w Walker) visitProperty(property *statement.Property) error {
	w.listener.EnterProperty(property)
	if property.PropertyValue != nil {
		err := w.visitExpression(property.PropertyValue)
		if err != nil {
			return err
		}
	}
	w.listener.ExitProperty(property)
	return nil
}

func (w Walker) visitPropertyAssignment(assignment *statement.PropertyAssignment) error {
	w.listener.EnterPropertyAssignment(assignment)
	err := w.visitExpression(assignment.Expression)
	if err != nil {
		return err
	}
	w.listener.ExitPropertyAssignment(assignment)
	return nil
}

func (w Walker) visitExpression(expr interface{}) error {
	switch e := expr.(type) {
	case *statement.ExplicitMemberExpression:
		w.listener.EnterExplicitMember(e)
		err := w.visitExpression(e.Expression)
		if err != nil {
			return err
		}
		w.listener.ExitExplicitMember(e)
	case *statement.FunctionCallExpression:
		err := w.visitFunctionCall(e)
		if err != nil {
			return err
		}
	case *literals.IntegerLiteral:
		w.listener.EnterInteger(e)
		w.listener.ExitInteger(e)
	case *literals.FloatingPointerLiteral:
		w.listener.EnterFloatPointer(e)
		w.listener.ExitFloatPointer(e)
	case *literals.StringLiteral:
		w.listener.EnterString(e)
		w.listener.ExitString(e)
	case *literals.BooleanLiteral:
		w.listener.EnterBoolean(e)
		w.listener.ExitBoolean(e)
	case literals.ListLiteral:
		w.listener.EnterList(e)
		iter := e.Front()
		for iter != nil {
			err := w.visitExpression(iter.Value)
			if err != nil {
				return err
			}
			iter = iter.Next()
		}
		w.listener.ExitList(e)
	default:
		return fmt.Errorf("unexpected expression type")
	}
	return nil
}

func (w Walker) visitFunctionCall(expression *statement.FunctionCallExpression) error {
	w.listener.EnterFunctionCall(expression)
	err := w.visitExpression(expression.Expression)
	if err != nil {
		return err
	}
	err = w.visitArgumentList(expression.ArgumentList)
	if err != nil {
		return err
	}
	w.listener.ExitFunctionCall(expression)
	return nil
}

func (w Walker) visitArgumentList(list statement.FunctionCallArgumentList) error {
	w.listener.EnterArgumentList(list)
	for arg := range list {
		err := w.visitExpression(arg)
		if err != nil {
			return err
		}
	}
	w.listener.ExitArgumentList(list)
	return nil
}

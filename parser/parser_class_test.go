package parser_test

import (
	"github.com/MetalheadSanya/VmlTranslator/parser"
	"github.com/MetalheadSanya/VmlTranslator/parser/literals"
	"github.com/MetalheadSanya/VmlTranslator/parser/statement"
	"strings"
	"testing"
)

func TestParserClassLine(t *testing.T) {
	str := "Rect { }"
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.Class
	if classStmt.Name != "Rect" {
		t.Errorf("Incorrect class name, expected Rect, take %q", classStmt.Name)
		return
	}
}

func TestParserClass(t *testing.T) {
	str := `Text {
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.Class
	if classStmt.Name != "Text" {
		t.Errorf("Incorrect class name, expected Text, take %q", classStmt.Name)
		return
	}
}

func TestParserClassPropertyAssignment(t *testing.T) {
	str := `
Rect {
    width: 100
    height: 100
    color: "red"
	enabled: false
	clip: true	
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	class := stmt.Class
	if class.Name != "Rect" {
		t.Errorf("Incorrect class class parent, expected 'Rectangle', take %s", class.Name)
	}

	checkPropertyAssignment(t, class)
}

func TestParserClassPropertyAssignmentSingleLine(t *testing.T) {
	str := `
Rectangle { width: 100; height: 100; color: "red"; enabled: false; clip: true }`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	class := stmt.Class
	if class.Name != "Rectangle" {
		t.Errorf("Incorrect class class parent, expected 'Rectangle', take %s", class.Name)
	}

	checkPropertyAssignment(t, class)
}

func checkPropertyAssignment(t *testing.T, class statement.Class) {
	propertyAssignmentCount := len(class.PropertyAssignments)
	if propertyAssignmentCount != 5 {
		t.Errorf("Incorrect property assignments count, expected '5', take %d", propertyAssignmentCount)
	}

	if propertyAssignmentCount < 1 {
		return
	}
	prop1 := class.PropertyAssignments[0]
	if prop1.Property.Len() != 1 {
		t.Errorf("Incorrect first property length")
	} else if *prop1.Property.Front().Value.(*string) != "width" {
		t.Errorf("Incorrect first property, expected 'width', take %s", *prop1.Property.Front().Value.(*string))
	}
	arg1, ok := prop1.Expression.(*literals.IntegerLiteral)
	if ok && *arg1 != 100 {
		t.Errorf("Incorrect first property value, expected '100', take %d", *arg1)
	} else if !ok {
		t.Errorf("Incorrect first property type, expected IntegerLiteral")
	}

	if propertyAssignmentCount < 2 {
		return
	}
	prop2 := class.PropertyAssignments[1]
	if prop2.Property.Len() != 1 {
		t.Errorf("Incorrect second property length")
	} else if *prop2.Property.Front().Value.(*string) != "height" {
		t.Errorf("Incorrect second property, expected 'height', take %s", *prop2.Property.Front().Value.(*string))
	}
	arg2, ok := prop2.Expression.(*literals.IntegerLiteral)
	if ok && *arg2 != 100 {
		t.Errorf("Incorrect second property value, expected '100', take %d", *arg2)
	} else if !ok {
		t.Errorf("Incorrect second property type, expected IntegerLiteral")
	}

	if propertyAssignmentCount < 3 {
		return
	}
	prop3 := class.PropertyAssignments[2]
	if prop3.Property.Len() != 1 {
		t.Errorf("Incorrect third property length")
	} else if *prop3.Property.Front().Value.(*string) != "color" {
		t.Errorf("Incorrect third property, expected 'color', take %s", *prop3.Property.Front().Value.(*string))
	}
	arg3, ok := prop3.Expression.(*literals.StringLiteral)
	if ok && *arg3 != "red" {
		t.Errorf("Incorrect third property value, expected 'red', take %s", *arg3)
	} else if !ok {
		t.Errorf("Incorrect third property type, expected StringLiteral")
	}

	if propertyAssignmentCount < 4 {
		return
	}
	prop4 := class.PropertyAssignments[3]
	if prop4.Property.Len() != 1 {
		t.Errorf("Incorrect fourth property length")
	} else if *prop4.Property.Front().Value.(*string) != "enabled" {
		t.Errorf("Incorrect fourth property, expected 'enabled', take %s", *prop4.Property.Front().Value.(*string))
	}
	arg4, ok := prop4.Expression.(*literals.BooleanLiteral)
	if ok && *arg4 != false {
		t.Errorf("Incorrect fourth property value, expected 'false', take %t", *arg4)
	} else if !ok {
		t.Errorf("Incorrect fourth property type, expected BooleanLiteral")
	}

	if propertyAssignmentCount < 5 {
		return
	}
	prop5 := class.PropertyAssignments[4]
	if prop5.Property.Len() != 1 {
		t.Errorf("Incorrect fifth property length")
	} else if *prop5.Property.Front().Value.(*string) != "clip" {
		t.Errorf("Incorrect fifth property, expected 'color', take %s", *prop5.Property.Front().Value.(*string))
	}
	arg5, ok := prop5.Expression.(*literals.BooleanLiteral)
	if ok && *arg5 != true {
		t.Errorf("Incorrect fifth property value, expected 'true', take %t", *arg5)
	} else if !ok {
		t.Errorf("Incorrect fifth property type, expected BooleanLiteral")
	}
}

func TestParserClassChildren(t *testing.T) {
	str := `
Rectangle {
	Rect { }
	Item { }
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	class := stmt.Class
	if class.Name != "Rectangle" {
		t.Errorf("Incorrect class class parent, expected 'Rectangle', take %s", class.Name)
	}

	checkChildren(t, class)
}

func checkChildren(t *testing.T, class statement.Class) {
	if class.Children.Len() != 2 {
		t.Errorf("expected 2 Children")
	}

	iter := class.Children.Front()
	if iter == nil {
		return
	}

	child1, ok := iter.Value.(*statement.Object)
	if !ok {
		t.Errorf("expected statement.Object")
	} else {
		if child1.Name.Len() != 1 {
			t.Errorf("expected first child object name be Rect, take Empty")
		} else if *child1.Name.Front().Value.(*string) != "Rect" {
			t.Errorf("expected first child object name Rect, take %q", *child1.Name.Front().Value.(*string))
		}
		if child1.Children.Len() != 0 {
			t.Errorf("exprected empty Children")
		}
		if child1.PropertyAssignments.Len() != 0 {
			t.Errorf("expected emptyPropertyAssignments")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}

	child2, ok := iter.Value.(*statement.Object)
	if !ok {
		t.Errorf("expected statement.Object")
	} else {
		if child2.Name.Len() != 1 {
			t.Errorf("expected second child object name be Rect, take Empty")
		} else if *child2.Name.Front().Value.(*string) != "Item" {
			t.Errorf("expected second child object name Rect, take %q", *child2.Name.Front().Value.(*string))
		}
		if child2.Children.Len() != 0 {
			t.Errorf("exprected empty Children")
		}
		if child2.PropertyAssignments.Len() != 0 {
			t.Errorf("expected emptyPropertyAssignments ")
		}
	}
}

func TestParserClassProperties(t *testing.T) {
	str := `
Rect {
	property color previousColor
	property double someNumber: 1.5
	default property string someString: "abc"
	property bool someBool: true
	property list<double> someNumber: [1.5, 1.0]
	property list<list<double>> someNumber
}`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.Class
	if classStmt.Name != "Rect" {
		t.Errorf("Incorrect class name, expected Rect, take %q", classStmt.Name)
		return
	}

	checkProperties(t, classStmt)
}

func checkProperties(t *testing.T, object statement.Class) {
	if object.Properties.Len() != 6 {
		t.Errorf("expected 6 properties, take %d", object.Properties.Len())
	}

	iter := object.Properties.Front()
	if iter == nil {
		return
	}
	stmt := iter.Value.(*statement.Property)

	if stmt.IsDefault {
		t.Errorf("Incorrect default value, expected false")
	}
	propertyName, ok := stmt.PropertyType.(*string)
	if ok && *propertyName != "color" {
		t.Errorf("Incorrect property type, expected 'color', take %q", stmt.PropertyType)
	} else if !ok {
		t.Errorf("Incorrect \"property type\" type, expected String")
	}
	if stmt.PropertyName != "previousColor" {
		t.Errorf("Incorrect property type, expected 'previousColor', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue != nil {
		t.Errorf("Incorrect property value, expected nil")
	}

	iter = iter.Next()
	if iter == nil {
		return
	}
	stmt = iter.Value.(*statement.Property)

	if stmt.IsDefault {
		t.Errorf("Incorrect default value, expected false")
	}
	propertyName, ok = stmt.PropertyType.(*string)
	if ok && *propertyName != "double" {
		t.Errorf("Incorrect property type, expected 'double', take %q", stmt.PropertyType)
	} else if !ok {
		t.Errorf("Incorrect \"property type\" type, expected String")
	}
	if stmt.PropertyName != "someNumber" {
		t.Errorf("Incorrect property type, expected 'someNumber', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue == nil {
		t.Errorf("Incorrect property value, expected FloatingPointerLiteral")
	} else {
		floatingPointerLiteral, ok := stmt.PropertyValue.(*literals.FloatingPointerLiteral)
		if ok && *floatingPointerLiteral != 1.5 {
			t.Errorf("Incorrect property value, expected '1.5', take %f", *floatingPointerLiteral)
		} else if !ok {
			t.Errorf("Incorrect property value, expected FloatingPointerLiteral")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}
	stmt = iter.Value.(*statement.Property)

	if !stmt.IsDefault {
		t.Errorf("Incorrect default value, expected false")
	}
	propertyName, ok = stmt.PropertyType.(*string)
	if ok && *propertyName != "string" {
		t.Errorf("Incorrect property type, expected 'string', take %q", stmt.PropertyType)
	} else if !ok {
		t.Errorf("Incorrect \"property type\" type, expected String")
	}
	if stmt.PropertyName != "someString" {
		t.Errorf("Incorrect property type, expected 'someString', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue == nil {
		t.Errorf("Incorrect property value, expected StringLiteral")
	} else {
		stringLiteral, ok := stmt.PropertyValue.(*literals.StringLiteral)
		if ok && *stringLiteral != "abc" {
			t.Errorf("Incorrect property value, expected 'abc', take %q", *stringLiteral)
		} else if !ok {
			t.Errorf("Incorrect property value, expected StringLiteral")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}
	stmt = iter.Value.(*statement.Property)

	if stmt.IsDefault {
		t.Errorf("Incorrect default value, expected false")
	}
	propertyName, ok = stmt.PropertyType.(*string)
	if ok && *propertyName != "bool" {
		t.Errorf("Incorrect property type, expected 'bool', take %q", stmt.PropertyType)
	} else if !ok {
		t.Errorf("Incorrect \"property type\" type, expected String")
	}
	if stmt.PropertyName != "someBool" {
		t.Errorf("Incorrect property type, expected 'someBool', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue == nil {
		t.Errorf("Incorrect property value, expected BooleanLiteral")
	} else {
		boolLiteral, ok := stmt.PropertyValue.(*literals.BooleanLiteral)
		if ok && *boolLiteral != true {
			t.Errorf("Incorrect property value, expected 'true', take %t", *boolLiteral)
		} else if !ok {
			t.Errorf("Incorrect property value, expected BooleanLiteral")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}
	stmt = iter.Value.(*statement.Property)

	if stmt.IsDefault {
		t.Errorf("Incorrect default value, expected false")
	}
	property, ok := stmt.PropertyType.(*statement.GenericType)
	if ok {
		if property.Container != "list" {
			t.Errorf("Incorrect container, expected 'list', take %q", property.Container)
		}
		element, ok := property.Element.(*string)
		if ok && *element != "double" {
			t.Errorf("Incorrect property type, expected 'double', take %q", *element)
		} else if !ok {
			t.Errorf("Incorrect element type, expected Identifier")
		}
	} else {
		t.Errorf("Incorrect \"property type\" type, expected statement.GenericType")
	}
	if stmt.PropertyName != "someNumber" {
		t.Errorf("Incorrect property type, expected 'someNumber', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue == nil {
		t.Errorf("Incorrect property value, expected ArrayLiteral")
	} else {
		listLiteral, ok := stmt.PropertyValue.(literals.ListLiteral)
		if ok {
			elementsCount := listLiteral.Len()
			if elementsCount != 2 {
				t.Errorf("Incorrect elements count, expected '2', take %d", elementsCount)
			}

			element := listLiteral.Front()
			if element == nil {
				return
			}
			el1, ok := element.Value.(*literals.FloatingPointerLiteral)
			if ok && *el1 != 1.5 {
				t.Errorf("Incorrect first property value, expected '1.5', take %f", *el1)
			} else if !ok {
				t.Errorf("Incorrect first property type, expected FloatingPointerLiteral")
			}

			element = element.Next()
			if element == nil {
				return
			}
			el2, ok := element.Value.(*literals.FloatingPointerLiteral)
			if ok && *el2 != 1 {
				t.Errorf("Incorrect second property value, expected '1', take %f", *el2)
			} else if !ok {
				t.Errorf("Incorrect second property type, expected FloatingPointerLiteral")
			}
		} else {
			t.Errorf("Incorrect property value, expected ArrayLiteral")
		}
	}

	iter = iter.Next()
	if iter == nil {
		return
	}
	stmt = iter.Value.(*statement.Property)

	property, ok = stmt.PropertyType.(*statement.GenericType)
	if ok {
		if property.Container != "list" {
			t.Errorf("Incorrect container, expected 'list', take %q", property.Container)
		}
		firstElement, ok := property.Element.(*statement.GenericType)
		if ok {
			if firstElement.Container != "list" {
				t.Errorf("Incorrect container, expected 'list', take %q", property.Container)
			}
			element, ok := firstElement.Element.(*string)
			if ok && *element != "double" {
				t.Errorf("Incorrect property type, expected 'double', take %q", *element)
			} else if !ok {
				t.Errorf("Incorrect element type, expected Identifier")
			}
		} else {
			t.Errorf("Incorrect \"property type\" type, expected statement.GenericType")
		}
	} else {
		t.Errorf("Incorrect \"property type\" type, expected statement.GenericType")
	}
	if stmt.PropertyName != "someNumber" {
		t.Errorf("Incorrect property type, expected 'someNumber', take %q", stmt.PropertyName)
	}
	if stmt.PropertyValue != nil {
		t.Errorf("Incorrect property value, expected nil")
	}
}

func TestClassFull(t *testing.T) {
	str := `
Rect {
	property color previousColor
    width: 100
	property double someNumber: 1.5
    height: 100
	default property string someString: "abc"

	Rect { }

	property bool someBool: true
    color: "red"

	Item { }
	property list<double> someNumber: [1.5, 1.0]

	enabled: false
	clip: true
	property list<list<double>> someNumber
}
`
	r := strings.NewReader(str)
	p := parser.VmlParser(r)
	stmt, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	classStmt := stmt.Class
	if classStmt.Name != "Rect" {
		t.Errorf("Incorrect class name, expected Rect, take %q", classStmt.Name)
		return
	}

	checkProperties(t, classStmt)
	checkChildren(t, classStmt)
	checkPropertyAssignment(t, classStmt)
}

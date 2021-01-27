package lispingo

import (
	"fmt"
	"strings"
)

func elementsToString(elements []Element) string {
	elementStrings := []string{}
	for _, e := range elements {
		elementStrings = append(elementStrings, e.toString())
	}
	return fmt.Sprintf("[%v]", strings.Join(elementStrings, ", "))
}

// Element is the interface for elements allowed in a list
type Element interface{
	toString()	string
}

// Int is the type for int literals
type Int struct {
	value   int
	lineNum int
}

func (i *Int) toString() string {
	return fmt.Sprintf("Int(%v, %v)", i.value, i.lineNum)
}

var _ Element = (*Int)(nil)

// String is the type for string literals
type String struct {
	value	string
	lineNum	int
}

func (s *String) toString() string {
	return fmt.Sprintf("String(%v, %v)", s.value, s.lineNum)
}
var _ Element = (*String)(nil)

// Symbol is the type for symbols
type Symbol struct {
	name	string
	lineNum	int
}

func (s *Symbol) toString() string {
	return fmt.Sprintf("Symbol(%v, %v)", s.name, s.lineNum)
}
var _ Element = (*Symbol)(nil)

// List is the type for list entities
type List struct {
	elements	[]Element	
	lineNum		int
}

func (l *List) toString() string {
	return fmt.Sprintf("List(%v, %v)", elementsToString(l.elements), l.lineNum)
}

var _ Element = (*List)(nil)

// FunctionCall is the type for function calls
type FunctionCall struct {
	function	Element
	arguments	[]Element
	lineNum		int
}

func (f *FunctionCall) toString() string {
	return fmt.Sprintf(
		"FunctionCall(%v, %v, %v)",
		f.function.toString(),
		elementsToString(f.arguments),
		f.lineNum,
	)
}

// ToString is the serializing method of FunctionCall for debugging
func (f *FunctionCall) ToString() string {
	return f.toString()
}
var _ Element = (*FunctionCall)(nil)

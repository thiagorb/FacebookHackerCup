package mrx

import "fmt"

type Expression interface {
	_getReduced() *Symbol
}

type Symbol struct {
	Value byte
}

type OperationStruct struct {
	Left  Expression
	Right Expression
}

type Operation interface {
	SetRight(e Expression)
	SetLeft(e Expression)
}

type OrOperation struct {
	OperationStruct
}

type AndOperation struct {
	OperationStruct
}

type XorOperation struct {
	OperationStruct
}

func (operation *OperationStruct) SetRight(e Expression) {
	operation.Right = e
}

func (operation *OperationStruct) SetLeft(e Expression) {
	operation.Left = e
}

func (symbol Symbol) String() string {
	return string(symbol.Value)
}

func (operation OrOperation) String() string {
	return fmt.Sprintf("(%s | %s)", operation.Left, operation.Right)
}

func (operation AndOperation) String() string {
	return fmt.Sprintf("(%s & %s)", operation.Left, operation.Right)
}

func (operation XorOperation) String() string {
	return fmt.Sprintf("(%s ^ %s)", operation.Left, operation.Right)
}

func (symbol *Symbol) _getReduced() *Symbol {
	return symbol
}

func (operation OrOperation) _getReduced() *Symbol {
	if GetReduced(&operation.Left) == Symbol1 || GetReduced(&operation.Right) == Symbol1 {
		return Symbol1
	}

	if GetReduced(&operation.Right) == Symbol0 {
		return GetReduced(&operation.Left)
	}

	if GetReduced(&operation.Left) == Symbol0 || GetReduced(&operation.Right) == GetReduced(&operation.Left) {
		return GetReduced(&operation.Right)
	}

	return Symbol1
}

func (operation AndOperation) _getReduced() *Symbol {
	if GetReduced(&operation.Left) == Symbol0 || GetReduced(&operation.Right) == Symbol0 {
		return Symbol0
	}

	if GetReduced(&operation.Right) == Symbol1 {
		return GetReduced(&operation.Left)
	}

	if GetReduced(&operation.Left) == Symbol1 || GetReduced(&operation.Right) == GetReduced(&operation.Left) {
		return GetReduced(&operation.Right)
	}

	return Symbol0
}

func (operation XorOperation) _getReduced() *Symbol {
	if GetReduced(&operation.Left) == Symbol0 {
		return GetReduced(&operation.Right)
	}

	if GetReduced(&operation.Right) == Symbol0 {
		return GetReduced(&operation.Left)
	}

	if GetReduced(&operation.Left) == Symbol1 {
		return Negate(GetReduced(&operation.Right))
	}

	if GetReduced(&operation.Right) == Symbol1 {
		return Negate(GetReduced(&operation.Left))
	}

	if GetReduced(&operation.Right) == GetReduced(&operation.Left) {
		return Symbol0
	}

	return Symbol1
}

var memReduced = map[Expression]*Symbol{}

func GetReduced(e *Expression) *Symbol {
	if _, ok := memReduced[*e]; !ok {
		memReduced[*e] = (*e)._getReduced()
	}

	return memReduced[*e]
}

func makeSymbol(value byte) *Symbol {
	s := new(Symbol)
	s.Value = value
	return s
}

var Symbol0 = makeSymbol('0')
var Symbol1 = makeSymbol('1')
var Symbolx = makeSymbol('x')
var SymbolX = makeSymbol('X')
var Symbols = [...]*Symbol{
	Symbol0,
	Symbol1,
	Symbolx,
	SymbolX,
}

func Negate(s *Symbol) *Symbol {
	switch s {
	case Symbol0:
		return Symbol1
	case Symbol1:
		return Symbol0
	case Symbolx:
		return SymbolX
	case SymbolX:
		return Symbolx
	}
	return nil
}

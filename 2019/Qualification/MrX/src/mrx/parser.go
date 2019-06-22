package mrx

import (
	"container/list"
)

func Parse(input string) Expression {
	stack := list.New()

	currentChar := 0

	for currentChar < len(input) {
		switch input[currentChar] {
		case 'x':
			stack.PushFront(Symbolx)
		case 'X':
			stack.PushFront(SymbolX)
		case '0':
			stack.PushFront(Symbol0)
		case '1':
			stack.PushFront(Symbol1)
		case '|':
			operation := new(OrOperation)
			operation.SetLeft(stack.Front().Value.(Expression))
			stack.Remove(stack.Front())
			stack.PushFront(operation)
		case '&':
			operation := new(AndOperation)
			operation.SetLeft(stack.Front().Value.(Expression))
			stack.Remove(stack.Front())
			stack.PushFront(operation)
		case '^':
			operation := new(XorOperation)
			operation.SetLeft(stack.Front().Value.(Expression))
			stack.Remove(stack.Front())
			stack.PushFront(operation)
		case ')':
			right := stack.Front()
			stack.Remove(right)
			operation := stack.Front().Value.(Operation)
			operation.SetRight(right.Value.(Expression))
		}
		currentChar++
	}

	return stack.Front().Value.(Expression)
}

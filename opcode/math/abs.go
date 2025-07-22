package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// AbsIdentifier defines the operation code (OpCode) for the absolute value operation in the virtual machine.
const AbsIdentifier opcode.OpCode = 12

// OpAbs defines a type representing an operation for computing the absolute value of a numeric value from the stack.
type OpAbs struct{}

// OpCode returns the operation code (OpCode) associated with the absolute value operation (AbsIdentifier).
func (o OpAbs) OpCode() opcode.OpCode {
	return AbsIdentifier
}

// Exec executes the absolute value operation on the top element of the stack and updates the execution context.
// Returns the program counter or an error if the stack is empty or the element is not a float64.
func (o OpAbs) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	n := ctx.Stack.Pop()

	if n == nil {
		return ctx.PC, fmt.Errorf("abs: stack underflow")
	}

	num, ok := n.(types.SymboliaNumber)
	if !ok {
		return ctx.PC, fmt.Errorf("abs: argument is not float64")
	}

	if num < 0 {
		ctx.Stack.Push(-num)
	} else {
		ctx.Stack.Push(num)
	}

	return 0, nil
}

// Name returns the name of the operation, which is "abs".
func (o OpAbs) Name() types.SymboliaString {
	return "abs"
}

package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// MinIdentifier defines the constant opcode representing the "min" operation in the instruction set.
const MinIdentifier = 23

// OpMin represents an operation that evaluates the minimum of two values from the execution stack.
type OpMin struct{}

// OpCode returns the opcode that identifies the "min" operation in the instruction set.
func (o OpMin) OpCode() opcode.OpCode {
	return MinIdentifier
}

// Exec executes the "min" operation, comparing and pushing the smaller value of two popped stack elements.
// Returns the new program counter and an error if a stack underflow occurs.
func (o OpMin) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	left := ctx.Stack.Pop()
	right := ctx.Stack.Pop()
	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("min: stack underflow")
	}
	if left.(types.SymboliaNumber) < right.(types.SymboliaNumber) {
		ctx.Stack.Push(left)
	} else {
		ctx.Stack.Push(right)
	}
	return ctx.PC + 1, nil
}

// Name returns the string representation of the operation, which is "min".
func (o OpMin) Name() types.SymboliaString {
	return "min"
}

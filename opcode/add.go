package opcode

import (
	"fmt"

	"github.com/riken127/symbolia/types"
)

// AddIdentifier is a constant representing the opcode for the addition operation in the virtual machine.
const AddIdentifier OpCode = 2

// OpAdd represents the addition operation in the virtual machine instruction set.
type OpAdd struct{}

// OpCode returns the opcode identifier for the addition operation.
func (o OpAdd) OpCode() OpCode { return AddIdentifier }

// Exec executes the addition operation on two float64 values popped from the stack and pushes the result back onto the stack.
// Returns the next program counter value or an error if types are invalid for the addition.
func (o OpAdd) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	b := ctx.Stack.Pop()
	a := ctx.Stack.Pop()
	af, aok := a.(float64)
	bf, bok := b.(float64)

	if !aok || !bok {
		return 0, fmt.Errorf("invalid types for addition: %T, %T", a, b)
	}

	ctx.Stack.Push(af + bf)
	return ctx.PC + 1, nil
}

// Name returns the human-readable name of the addition operation.
func (o OpAdd) Name() string {
	return "add"
}

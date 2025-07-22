package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

// MaxIdentifier is a constant representing the operation code for the "max" instruction in the execution context.
const MaxIdentifier = 22

// OpMax is a type representing the "max" operation, used to determine the greater of two values from the stack.
type OpMax struct{}

// OpCode returns the operation code identifier for the OpMax operation.
func (o OpMax) OpCode() OpCode {
	return MaxIdentifier
}

// Exec executes the "max" operation, comparing the top two stack values and pushing the greater one back on the stack.
// Returns the next program counter value or an error if the stack has insufficient values.
func (o OpMax) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	left := ctx.Stack.Pop()
	right := ctx.Stack.Pop()
	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("max: stack underflow")
	}
	if left.(float64) > right.(float64) {
		ctx.Stack.Push(left)
	} else {
		ctx.Stack.Push(right)
	}

	return ctx.PC + 1, nil
}

// Name returns the name of the operation, which is "max".
func (o OpMax) Name() string {
	return "max"
}

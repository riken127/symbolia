package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

// GeIdentifier represents the opcode for the "greater than or equal" comparison operation.
const GeIdentifier OpCode = 17

// OpGe represents the "greater than or equal" comparison operation in the execution context.
type OpGe struct{}

// OpCode returns the operation code associated with the "greater than or equal" comparison operation.
func (o OpGe) OpCode() OpCode {
	return GeIdentifier
}

// Exec evaluates the "greater than or equal" operation, compares two numeric values, and pushes the result on the stack.
// Returns the next program counter value or an error if the operation fails.
func (o OpGe) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("ge: stack underflow")
	}

	rightFloat, ok1 := right.(float64)
	leftFloat, ok2 := left.(float64)

	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("ge: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftFloat >= rightFloat)
	return ctx.PC + 1, nil
}

// Name returns the identifier name of the "greater than or equal" operation as a string.
func (o OpGe) Name() string {
	return "ge"
}

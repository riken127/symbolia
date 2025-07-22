package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

// AndIdentifier is the OpCode representing the logical "AND" operation in instruction execution.
const AndIdentifier OpCode = 13

// OpAnd represents the logical "AND" operation to be executed as an instruction within the execution context.
type OpAnd struct{}

// OpCode returns the operation code associated with the logical "AND" instruction.
func (o OpAnd) OpCode() OpCode { return AndIdentifier }

// Exec executes the "AND" operation using operands from the instruction arguments or stack, returning the updated PC or an error.
func (o OpAnd) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	var left, right any

	if len(instr.Args) == 2 {
		left, right = instr.Args[0], instr.Args[1]
	} else if len(instr.Args) == 1 {
		right = instr.Args[0]
		left = ctx.Stack.Pop()
		if left == nil {
			return ctx.PC, fmt.Errorf("and: stack underflow")
		}
	} else {
		right = ctx.Stack.Pop()
		left = ctx.Stack.Pop()
		if left == nil || right == nil {
			return ctx.PC, fmt.Errorf("and: stack underflow")
		}
	}

	leftBool, ok1 := left.(bool)
	rightBool, ok2 := right.(bool)
	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("and: operands must be boolean values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftBool && rightBool)
	return ctx.PC + 1, nil
}

// Name returns the string representation of the logical "AND" operation.
func (o OpAnd) Name() string { return "and" }

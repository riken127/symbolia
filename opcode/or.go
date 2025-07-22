package opcode

import (
	"fmt"
	"github.com/riken127/symbolia/types"
)

// OrIdentifier represents the opcode for the logical OR operation in the instruction set.
const OrIdentifier OpCode = 14

// OpOr represents the logical OR operation to be executed during interpretation or execution of instructions.
type OpOr struct{}

// OpCode returns the operation code associated with the logical OR operation.
func (o OpOr) OpCode() OpCode { return OrIdentifier }

// Exec executes the logical OR operation using the provided execution context and instruction's arguments or stack.
// Returns the updated program counter or an error if operands are invalid or stack underflows.
func (o OpOr) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	var left, right any

	if len(instr.Args) == 2 {
		left, right = instr.Args[0], instr.Args[1]
	} else if len(instr.Args) == 1 {
		right = instr.Args[0]
		left = ctx.Stack.Pop()
		if left == nil {
			return ctx.PC, fmt.Errorf("or: stack underflow")
		}
	} else {
		right = ctx.Stack.Pop()
		left = ctx.Stack.Pop()
		if left == nil || right == nil {
			return ctx.PC, fmt.Errorf("or: stack underflow")
		}
	}

	leftBool, ok1 := left.(bool)
	rightBool, ok2 := right.(bool)
	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("or: operands must be boolean values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftBool || rightBool)
	return ctx.PC + 1, nil
}

// Name returns the name of the logical OR operation as a string.
func (o OpOr) Name() string { return "or" }

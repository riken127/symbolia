package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"

	"github.com/riken127/symbolia/types"
)

// EqualIdentifier is a constant representing the opcode for the equality comparison operation in the virtual machine.
const EqualIdentifier = 10

// OpEqual represents the equality comparison operation in the virtual machine instruction set.
type OpEqual struct{}

// OpCode returns the operation code corresponding to the equality comparison operation.
func (o OpEqual) OpCode() opcode.OpCode {
	return EqualIdentifier
}

// Name returns the name of the operation as a string identifier.
func (o OpEqual) Name() types.SymboliaString {
	return "equal"
}

// Exec performs the equality comparison operation by popping two values from the stack, comparing them, and pushing the result.
// Returns the updated program counter or an error if the stack does not contain enough elements.
func (o OpEqual) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if right == nil || left == nil {
		return ctx.PC, fmt.Errorf("equal: stack underflow")
	}
	ctx.Stack.Push(left == right)
	return ctx.PC + 1, nil
}

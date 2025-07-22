package misc

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// RotIdentifier represents the opcode for the `rot` operation
const RotIdentifier opcode.OpCode = 29

// OpRot represents the operation that rotates the top 3 elements of the stack
type OpRot struct{}

// OpCode returns the unique opcode for the `rot` operation
func (o OpRot) OpCode() opcode.OpCode {
	return RotIdentifier
}

// Name returns the human-readable name of the operation
func (o OpRot) Name() types.SymboliaString {
	return "rot"
}

// Exec rotates the top 3 elements of the stack
// - The third value from the top becomes the top value
// - The second value remains in the middle
// - The top value moves to the third position
// Returns an error if there are fewer than 3 elements
func (o OpRot) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	// Pop the top 3 elements
	c := ctx.Stack.Pop()
	b := ctx.Stack.Pop()
	a := ctx.Stack.Pop()

	// Check if any element is nil (stack underflow)
	if a == nil || b == nil || c == nil {
		return ctx.PC, fmt.Errorf("rot: stack underflow, requires at least 3 elements")
	}

	// Push them back in rotated order
	ctx.Stack.Push(b) // Middle value now remains in the middle
	ctx.Stack.Push(c) // Top value becomes the second
	ctx.Stack.Push(a) // Third value becomes the top

	return ctx.PC + 1, nil
}

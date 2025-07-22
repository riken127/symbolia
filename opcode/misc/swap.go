package misc

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// SwapIdentifier represents the opcode for the `swap` operation
const SwapIdentifier opcode.OpCode = 30

// OpSwap represents the operation to swap the top two elements of the stack
type OpSwap struct{}

// OpCode returns the unique opcode for the `swap` operation
func (o OpSwap) OpCode() opcode.OpCode {
	return SwapIdentifier
}

// Name returns the human-readable name of the operation
func (o OpSwap) Name() types.SymboliaString {
	return "swap"
}

// Exec swaps the top two elements of the stack
// - The top element becomes the second
// - The second element becomes the top
// Returns an error if there are fewer than 2 elements
func (o OpSwap) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	// Pop the top 2 elements
	top := ctx.Stack.Pop()
	second := ctx.Stack.Pop()

	// Check if any element is nil (stack underflow)
	if top == nil || second == nil {
		return ctx.PC, fmt.Errorf("swap: stack underflow, requires at least 2 elements")
	}

	// Push them back in swapped order
	ctx.Stack.Push(top)    // Second element becomes the top
	ctx.Stack.Push(second) // Top element becomes the second

	return ctx.PC + 1, nil
}

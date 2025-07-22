package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// ModIdentifier is the opcode representation for the modulus operation in the instruction set.
const ModIdentifier = 24

// OpMod is a type representing the modulus operation in the instruction set.
type OpMod struct{}

// OpCode returns the opcode identifier for the modulus operation.
func (o OpMod) OpCode() opcode.OpCode {
	return ModIdentifier
}

// Exec performs the modulus operation with the top two stack values and pushes the result back onto the stack.
// Returns the updated program counter and an error if operands are invalid or if there is a stack underflow.
func (o OpMod) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	left := ctx.Stack.Pop()
	right := ctx.Stack.Pop()
	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("mod: stack underflow")
	}

	leftFloat, ok1 := left.(types.SymboliaInteger)
	rightFloat, ok2 := right.(types.SymboliaInteger)
	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("mod: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(((rightFloat % leftFloat) + leftFloat) % leftFloat)

	return ctx.PC + 1, nil
}

// Name returns the string identifier for the modulus operation.
func (o OpMod) Name() types.SymboliaString {
	return "mod"
}

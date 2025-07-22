package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// IncIdentifier represents the operation code for the increment instruction, used to increase a numeric value by one.
const IncIdentifier opcode.OpCode = 21

// OpInc represents an operation to increment the top numeric value on the stack by 1.
type OpInc struct{}

// OpCode returns the operation code associated with the increment instruction (IncIdentifier).
func (o OpInc) OpCode() opcode.OpCode {
	return IncIdentifier
}

// Exec increments the top numeric value on the stack by 1, updating the program counter, and returns any errors encountered.
func (o OpInc) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()
	if val == nil {
		return ctx.PC, fmt.Errorf("inc: stack underflow")
	}

	num, ok := val.(types.SymboliaNumber)
	if !ok {
		return ctx.PC, fmt.Errorf("inc: invalid type, expected a numeric value")
	}

	ctx.Stack.Push(num + 1)
	return ctx.PC + 1, nil
}

// Name returns the string representation of the operation, which is "inc".
func (o OpInc) Name() types.SymboliaString {
	return "inc"
}

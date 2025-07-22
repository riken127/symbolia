package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// DecIdentifier is the OpCode that represents the "dec" operation, which decrements a value on the stack by 1.
const DecIdentifier opcode.OpCode = 16

// OpDec represents an operation that decrements a value on the stack by 1.
type OpDec struct{}

// OpCode returns the operation code (OpCode) associated with the "dec" operation.
func (o OpDec) OpCode() opcode.OpCode {
	return DecIdentifier
}

// Exec executes the "dec" operation, decrementing the top numeric value on the stack by 1 and updating the program counter.
func (o OpDec) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()
	if val == nil {
		return ctx.PC, fmt.Errorf("dec: stack underflow")
	}

	num, ok := val.(types.SymboliaNumber)
	if !ok {
		return ctx.PC, fmt.Errorf("dec: invalid type, expected a numeric value")
	}

	ctx.Stack.Push(num - 1)
	return ctx.PC + 1, nil
}

// Name returns the name of the operation associated with OpDec, which is "dec".
func (o OpDec) Name() types.SymboliaString {
	return "dec"
}

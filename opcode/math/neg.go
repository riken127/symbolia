package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// NegIdentifier is a constant representing the operation code for the negation operation.
const NegIdentifier = 25

// OpNeg represents an operation that negates a numeric value from the stack and pushes the result back onto the stack.
type OpNeg struct{}

// OpCode returns the operation code associated with the OpNeg operation.
func (o OpNeg) OpCode() opcode.OpCode {
	return NegIdentifier
}

// Exec performs the negation operation on the top numeric value of the stack and pushes the result back onto the stack.
// Returns the updated program counter or an error if the stack is empty or the value is not numeric.
func (o OpNeg) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Pop()
	if val == nil {
		return ctx.PC, fmt.Errorf("neg: stack underflow")
	}

	num, ok := val.(types.SymboliaNumber)

	if !ok {
		return ctx.PC, fmt.Errorf("neg: invalid type, expected a numeric value")
	}

	ctx.Stack.Push(-num)

	return ctx.PC + 1, nil
}

// Name returns the string identifier for the OpNeg operation, which is "neg".
func (o OpNeg) Name() types.SymboliaString {
	return "neg"
}

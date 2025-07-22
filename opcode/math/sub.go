package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"

	"github.com/riken127/symbolia/types"
)

// SubIdentifier is a constant representing the opcode for the subtraction operation in the virtual machine.
const SubIdentifier opcode.OpCode = 3

// OpSub represents the subtraction operation in the virtual machine instruction set.
type OpSub struct{}

// OpCode returns the operation code associated with the subtraction operation.
func (o OpSub) OpCode() opcode.OpCode { return SubIdentifier }

// Exec performs the subtraction operation by popping two values from the stack, subtracting them, and pushing the result.
// Returns the updated program counter or an error if the popped values are not float64.
func (o OpSub) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	b := ctx.Stack.Pop()
	a := ctx.Stack.Pop()
	af, aok := a.(types.SymboliaNumber)
	bf, bok := b.(types.SymboliaNumber)
	if !aok || !bok {
		return 0, fmt.Errorf("sub: valores inválidos %T, %T", a, b)
	}
	ctx.Stack.Push(af - bf)
	return ctx.PC + 1, nil
}

// Name returns the name of the subtraction operation as a string.
func (o OpSub) Name() types.SymboliaString { return "sub" }

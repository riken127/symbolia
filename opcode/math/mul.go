package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"

	"github.com/riken127/symbolia/types"
)

// MulIdentifier is a constant representing the opcode for the multiplication operation in the virtual machine.
const MulIdentifier = 4

// OpMul represents the multiplication operation in the virtual machine instruction set.
type OpMul struct{}

// OpCode returns the operation code identifier for the multiplication operation.
func (o OpMul) OpCode() opcode.OpCode { return MulIdentifier }

// Exec executes the multiplication operation, popping two values from the stack, multiplying them, and pushing the result.
// Returns the next program counter and an error if the popped values are not valid float64 types.
func (o OpMul) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	b := ctx.Stack.Pop()
	a := ctx.Stack.Pop()
	af, aok := a.(types.SymboliaNumber)
	bf, bok := b.(types.SymboliaNumber)
	if !aok || !bok {
		return 0, fmt.Errorf("mul: valores inválidos %T, %T", a, b)
	}
	ctx.Stack.Push(af * bf)
	return ctx.PC + 1, nil
}

// Name returns the name identifier of the multiplication operation.
func (o OpMul) Name() types.SymboliaString { return "mul" }

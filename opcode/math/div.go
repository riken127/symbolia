package math

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"

	"github.com/riken127/symbolia/types"
)

// DivIdentifier is a constant representing the opcode for the division operation in the virtual machine.
const DivIdentifier = 5

// OpDiv represents the division operation in the virtual machine instruction set.
type OpDiv struct{}

// OpCode returns the operation code identifier for the division operation.
func (o OpDiv) OpCode() opcode.OpCode { return DivIdentifier }

// Exec performs the division operation, validates operands, and updates the program counter and execution stack.
func (o OpDiv) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	b := ctx.Stack.Pop()
	a := ctx.Stack.Pop()
	af, aok := a.(types.SymboliaNumber)
	bf, bok := b.(types.SymboliaNumber)
	if !aok || !bok {
		return 0, fmt.Errorf("invalid values for division: %T, %T", a, b)
	}
	if bf == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	ctx.Stack.Push(af / bf)
	return ctx.PC + 1, nil
}

// Name returns the name identifier for the division operation.
func (o OpDiv) Name() types.SymboliaString { return "div" }

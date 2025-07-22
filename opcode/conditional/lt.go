package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// LtIdentifier is the OpCode representing the "less than" operation in instructions.
const LtIdentifier opcode.OpCode = 20

// OpLt represents the "less than" operation, used to compare two numeric values on the execution stack.
type OpLt struct{}

// OpCode returns the operation code identifier for the OpLt operation.
func (o OpLt) OpCode() opcode.OpCode {
	return LtIdentifier
}

// Exec performs the "less than" operation, comparing two numeric values popped from the stack, and pushes the result.
// Returns the next program counter value and an error if operands are not numeric or stack is underflowed.
func (o OpLt) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("lt: stack underflow")
	}

	rightFloat, ok1 := right.(types.SymboliaNumber)
	leftFloat, ok2 := left.(types.SymboliaNumber)

	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("lt: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftFloat < rightFloat)
	return ctx.PC + 1, nil
}

// Name returns the string representation of the "less than" operation.
func (o OpLt) Name() types.SymboliaString {
	return "lt"
}

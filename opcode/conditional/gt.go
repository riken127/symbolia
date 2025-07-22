package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// GtIdentifier is the OpCode for the "greater than" operation, which checks if the left operand is greater than the right.
const GtIdentifier opcode.OpCode = 18

// OpGt represents the "greater than" operation, comparing two numeric values and pushing the result onto the stack.
type OpGt struct{}

// OpCode returns the operation code identifying the "greater than" operation.
func (o OpGt) OpCode() opcode.OpCode {
	return GtIdentifier
}

// Exec performs the "greater than" operation by comparing the top two numeric values from the stack and pushing the result.
// Returns the next program counter or an error if the stack is underflowed or operands are not numeric.
func (o OpGt) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("gt: stack underflow")
	}

	rightFloat, ok1 := right.(types.SymboliaNumber)
	leftFloat, ok2 := left.(types.SymboliaNumber)

	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("gt: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftFloat > rightFloat)
	return ctx.PC + 1, nil
}

// Name returns the name of the operation as a string, "gt".
func (o OpGt) Name() types.SymboliaString {
	return "gt"
}

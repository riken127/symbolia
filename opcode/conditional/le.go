package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// LeIdentifier represents the operation code for the "less than or equal" comparison instruction.
const LeIdentifier opcode.OpCode = 19

// OpLe represents the "less than or equal" operation used in instruction execution within a stack-based context.
type OpLe struct{}

// OpCode returns the operation code corresponding to the "less than or equal" instruction.
func (o OpLe) OpCode() opcode.OpCode {
	return LeIdentifier
}

// Exec executes the "less than or equal" operation, comparing two numeric operands from the stack.
// Returns the updated program counter or an error if the stack is underflowed or operands are non-numeric.
func (o OpLe) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("le: stack underflow")
	}

	rightFloat, ok1 := right.(types.SymboliaNumber)
	leftFloat, ok2 := left.(types.SymboliaNumber)

	if !ok1 || !ok2 {
		return ctx.PC, fmt.Errorf("le: operands must be numeric values (got %T, %T)", left, right)
	}

	ctx.Stack.Push(leftFloat <= rightFloat)
	return ctx.PC + 1, nil
}

// Name returns the string identifier of the "less than or equal" operation.
func (o OpLe) Name() types.SymboliaString {
	return "le"
}

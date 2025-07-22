package conditional

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// DiffIdentifier is a constant representing the opcode for the "nequal" operation, used to compare two stack values.
const DiffIdentifier = 26

// OpDiff represents an operation that performs a "not equal" comparison between two stack values.
type OpDiff struct{}

// OpCode returns the operation code associated with the OpDiff struct.
func (o OpDiff) OpCode() opcode.OpCode {
	return DiffIdentifier
}

// Exec performs the "not equal" comparison on the top two stack values and updates the program counter.
// Returns an error if there are insufficient values on the stack.
func (o OpDiff) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	right := ctx.Stack.Pop()
	left := ctx.Stack.Pop()

	if left == nil || right == nil {
		return ctx.PC, fmt.Errorf("nequal: stack underflow")
	}

	ctx.Stack.Push(left != right)

	return ctx.PC + 1, nil
}

// Name returns the string identifier for the OpDiff operation, which is "nequal".
func (o OpDiff) Name() types.SymboliaString {
	return "nequal"
}

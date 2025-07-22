package misc

import (
	"fmt"
	"github.com/riken127/symbolia/opcode"

	"github.com/riken127/symbolia/types"
)

// DupIdentifier is a constant representing the opcode for duplicating the top value of the stack in the virtual machine.
const DupIdentifier opcode.OpCode = 9

// OpDup represents the operation that duplicates the top element of the stack in a virtual machine.
type OpDup struct{}

// OpCode retrieves the operation code identifying the specific operation implemented by OpDup.
func (o OpDup) OpCode() opcode.OpCode { return DupIdentifier }

// Name returns the name of the operation implemented by OpDup as a string.
func (o OpDup) Name() types.SymboliaString { return "dup" }

// Exec performs the duplication of the top element of the stack in the execution context.
// Returns the updated program counter on success or an error if the stack is empty.
func (o OpDup) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	val := ctx.Stack.Peek()
	if val == nil {
		return 0, fmt.Errorf("dup: stack underflow")
	}
	ctx.Stack.Push(val)
	return ctx.PC + 1, nil
}

package misc

import (
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

// PushIdentifier is a constant representing the opcode for the push operation in the virtual machine.
const PushIdentifier opcode.OpCode = 1

// OpPush represents an instruction for pushing arguments onto the stack in a virtual machine.
type OpPush struct{}

// OpCode returns the operation code associated with the OpPush instruction.
func (o OpPush) OpCode() opcode.OpCode { return PushIdentifier }

// Exec executes the push operation, adding each argument from the instruction to the stack, and increments the program counter.
func (o OpPush) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	for _, arg := range instr.Args {
		ctx.Stack.Push(arg)
	}
	return ctx.PC + 1, nil
}

// Name returns the string representation of the OpPush instruction, typically used for identification or debugging.
func (o OpPush) Name() types.SymboliaString {
	return "push"
}

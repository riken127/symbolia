package opcode

import (
	"fmt"

	"github.com/riken127/symbolia/types"
)

// JumpIfIdentifier is a constant representing the opcode for conditional jumps based on a boolean value in the stack.
const JumpIfIdentifier OpCode = 8

// OpJumpIf represents a conditional jump operation that depends on a boolean value at the top of the execution stack.
type OpJumpIf struct{}

// OpCode returns the operation code representing the conditional jump for the OpJumpIf instruction.
func (o OpJumpIf) OpCode() OpCode { return JumpIfIdentifier }

// Exec evaluates the top value of the stack as a boolean to determine whether to jump to a target instruction.
// Returns the next program counter value or an error if the stack is underflowed, the condition is invalid, or the target is missing.
func (o OpJumpIf) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	cond := ctx.Stack.Pop()
	if cond == nil {
		return ctx.PC, fmt.Errorf("jumpif: stack underflow")
	}

	condBool, ok := cond.(bool)
	if !ok {
		return ctx.PC, fmt.Errorf("jumpif condition is not bool")
	}

	if condBool {
		if len(instr.Args) == 0 {
			return ctx.PC, fmt.Errorf("jumpif: missing jump target")
		}
		target, ok := instr.Args[0].(int)
		if !ok {
			return ctx.PC, fmt.Errorf("jumpif: jump target is not int")
		}
		return target, nil
	}

	return ctx.PC + 1, nil
}

// Name returns the name of the operation, "jumpif".
func (o OpJumpIf) Name() string { return "jumpif" }

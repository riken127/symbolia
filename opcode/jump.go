package opcode

import (
	"fmt"

	"github.com/riken127/symbolia/types"
)

// JumpIdentifier is a constant representing the opcode for the jump operation in the virtual machine.
const JumpIdentifier OpCode = 7

// OpJump represents the jump operation in the virtual machine instruction set.
type OpJump struct{}

// OpCode returns the operation code associated with the jump instruction.
func (o OpJump) OpCode() OpCode { return JumpIdentifier }

// Exec executes the jump operation, validating and extracting the jump target from the instruction arguments.
func (o OpJump) Exec(ctx *types.ExecutionContext, instr types.Instruction) (int, error) {
	if len(instr.Args) != 1 {
		return 0, fmt.Errorf("jump needs 1 argument")
	}

	target, ok := instr.Args[0].(int)
	if !ok {
		return 0, fmt.Errorf("jump argument is not int")
	}

	return target, nil
}

// Name returns the name of the operation as a string.
func (o OpJump) Name() string { return "jump" }

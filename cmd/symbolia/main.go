package main

import (
	"log"

	"github.com/riken127/symbolia/engine"
	"github.com/riken127/symbolia/opcode"
	"github.com/riken127/symbolia/types"
)

func main() {
	// Function 1 - increments the counter 'til 3 and passes the value via pipeTo
	fn1 := types.AbstractFunction{
		ID:      1,
		Symbols: []string{"counter"},
		Stack: []types.Instruction{
			{OpCode: int(opcode.PushIdentifier), Args: []any{"counter"}}, // 0: push counter
			{OpCode: int(opcode.PushIdentifier), Args: []any{3.0}},       // 1: push 3
			{OpCode: int(opcode.EqualIdentifier)},                        // 2: equal
			{OpCode: int(opcode.JumpIfIdentifier), Args: []any{9}},       // 3: jump if equal → pc=9

			{OpCode: int(opcode.PushIdentifier), Args: []any{"counter"}},  // 4: push counter
			{OpCode: int(opcode.PushIdentifier), Args: []any{1.0}},        // 5: push 1
			{OpCode: int(opcode.AddIdentifier)},                           // 6: add
			{OpCode: int(opcode.StoreIdentifier), Args: []any{"counter"}}, // 7: store in counter
			{OpCode: int(opcode.JumpIdentifier), Args: []any{0}},          // 8: jump to loop start

			{OpCode: int(opcode.PushIdentifier), Args: []any{"counter"}}, // 9: final push for pipe
		},
	}

	// Functions 2 - multiplies per 10 and prints
	fn2 := types.AbstractFunction{
		ID:      2,
		Symbols: []string{"value"},
		Stack: []types.Instruction{
			{OpCode: int(opcode.PushIdentifier), Args: []any{"value"}}, // get pipe'd value
			{OpCode: int(opcode.PushIdentifier), Args: []any{10.0}},
			{OpCode: int(opcode.MulIdentifier)},
			{OpCode: int(opcode.PrintIdentifier)},
		},
	}

	// Invoke 1st via PipeTo → 2
	inv := types.Invocation{
		FunctionID: 1,
		Args:       []any{0.0}, // counter starts at 0
		PipeTo:     2,          // pipe final result on the stack to function 2
	}

	// Executor with both functions loaded
	exec := engine.NewExecutor([]types.AbstractFunction{fn1, fn2}, inv)

	// Execute
	if err := exec.ExecuteInvocation(inv); err != nil {
		log.Fatalf("Error @: %v", err)
	}
}

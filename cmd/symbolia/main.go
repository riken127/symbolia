package main

import (
	"github.com/riken127/symbolia/opcode/conditional"
	"github.com/riken127/symbolia/opcode/flow"
	"github.com/riken127/symbolia/opcode/math"
	"github.com/riken127/symbolia/opcode/misc"
	"log"
	"time"

	"github.com/riken127/symbolia/engine"
	"github.com/riken127/symbolia/types"
)

func main() {
	// Função 1 - incrementa 'counter' até 1_000_000 e passa o valor via pipeTo
	fn1 := types.AbstractFunction{
		ID:      1,
		Symbols: []string{"counter"},
		Stack: []types.Instruction{
			// Loop start:
			{OpCode: int(misc.PushIdentifier), Args: []any{"counter"}},   // 0: push counter
			{OpCode: int(misc.PushIdentifier), Args: []any{1_000_000.0}}, // 1: push limit
			{OpCode: int(conditional.EqualIdentifier)},                   // 2: equal ?
			{OpCode: int(flow.JumpIfIdentifier), Args: []any{9}},         // 3: jump if equal → end loop

			{OpCode: int(misc.PushIdentifier), Args: []any{"counter"}},  // 4: push counter
			{OpCode: int(misc.PushIdentifier), Args: []any{1.0}},        // 5: push 1
			{OpCode: int(math.AddIdentifier)},                           // 6: counter + 1
			{OpCode: int(misc.StoreIdentifier), Args: []any{"counter"}}, // 7: store counter
			{OpCode: int(flow.JumpIdentifier), Args: []any{0}},          // 8: jump to loop start

			// End loop: push final counter value
			{OpCode: int(misc.PushIdentifier), Args: []any{"counter"}}, // 9: push counter for pipe
		},
	}

	// Função 2 - multiplica o valor por 10 e imprime
	fn2 := types.AbstractFunction{
		ID:      2,
		Symbols: []string{"value"},
		Stack: []types.Instruction{
			{OpCode: int(misc.PushIdentifier), Args: []any{"value"}}, // push value
			{OpCode: int(misc.PushIdentifier), Args: []any{10.0}},
			{OpCode: int(math.MulIdentifier)},
			{OpCode: int(misc.PrintIdentifier)},
		},
	}

	// Invocation que inicia o contador a 0 e pipe para a função 2
	inv := types.Invocation{
		FunctionID: 1,
		Args:       []any{0.0},
		PipeTo:     2,
	}

	exec := engine.NewExecutor([]types.AbstractFunction{fn1, fn2}, inv)

	start := time.Now()
	if err := exec.ExecuteInvocation(inv); err != nil {
		log.Fatalf("Execution error: %v", err)
	}
	elapsed := time.Since(start)

	log.Printf("\nBenchmark finished in %s", elapsed)
}

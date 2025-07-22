package opcode

import (
	"github.com/riken127/symbolia/opcode/conditional"
	"github.com/riken127/symbolia/opcode/flow"
	"github.com/riken127/symbolia/opcode/math"
	"github.com/riken127/symbolia/opcode/misc"
)

// Registry is a mapping of OpCode identifiers to their corresponding InstructionHandler implementations.
var Registry = map[OpCode]InstructionHandler{
	math.AddIdentifier:          math.OpAdd{},
	misc.PushIdentifier:         misc.OpPush{},
	math.SubIdentifier:          math.OpSub{},
	math.MulIdentifier:          math.OpMul{},
	math.DivIdentifier:          math.OpDiv{},
	misc.PrintIdentifier:        misc.OpPrint{},
	flow.JumpIdentifier:         flow.OpJump{},
	flow.JumpIfIdentifier:       flow.OpJumpIf{},
	misc.DupIdentifier:          misc.OpDup{},
	conditional.EqualIdentifier: conditional.OpEqual{},
	misc.StoreIdentifier:        misc.OpStore{},
	math.AbsIdentifier:          math.OpAbs{},
	conditional.AndIdentifier:   conditional.OpAnd{},
	misc.ClearIdentifier:        misc.OpClear{},
	math.DecIdentifier:          math.OpDec{},
	conditional.DiffIdentifier:  conditional.OpDiff{},
	conditional.GeIdentifier:    conditional.OpGe{},
	conditional.GtIdentifier:    conditional.OpGt{},
	math.IncIdentifier:          math.OpInc{},
	conditional.LeIdentifier:    conditional.OpLe{},
	conditional.LtIdentifier:    conditional.OpLt{},
	math.MaxIdentifier:          math.OpMax{},
	math.MinIdentifier:          math.OpMin{},
	math.ModIdentifier:          math.OpMod{},
	math.NegIdentifier:          math.OpNeg{},
	conditional.NotIdentifier:   conditional.OpNot{},
	conditional.OrIdentifier:    conditional.OpOr{},
	misc.PopIdentifier:          misc.OpPop{},
	misc.RotIdentifier:          misc.OpRot{},
	misc.SwapIdentifier:         misc.OpSwap{},
	misc.TypeIdentifier:         misc.OpType{},
}

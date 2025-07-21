package opcode

// OpCodeRegistry is a mapping of OpCode identifiers to their corresponding InstructionHandler implementations.
var OpCodeRegistry = map[OpCode]InstructionHandler{
	AddIdentifier:    OpAdd{},
	PushIdentifier:   OpPush{},
	SubIdentifier:    OpSub{},
	MulIdentifier:    OpMul{},
	DivIdentifier:    OpDiv{},
	PrintIdentifier:  OpPrint{},
	JumpIdentifier:   OpJump{},
	JumpIfIdentifier: OpJumpIf{},
	DupIdentifier:    OpDup{},
	EqualIdentifier:  OpEqual{},
	StoreIdentifier:  OpStore{},
}

package wax

type Opcode byte

const (
	// Control flow operators
	OpcodeUnreachable Opcode = 0x00
	OpcodeNop         Opcode = 0x01
	OpcodeBlock       Opcode = 0x02
	OpcodeLoop        Opcode = 0x03
	OpcodeIf          Opcode = 0x04
	OpcodeElse        Opcode = 0x05
	OpcodeEnd         Opcode = 0x0b
	OpcodeBr          Opcode = 0x0c
	OpcodeBrIf        Opcode = 0x0d
	OpcodeBrTable     Opcode = 0x0e
	OpcodeReturn      Opcode = 0x0f

	// Call operators
	OpcodeCall         Opcode = 0x10
	OpcodeCallIndirect Opcode = 0x11

	// Parametric operators
	OpcodeDrop   Opcode = 0x1a
	OpcodeSelect Opcode = 0x1b

	// Variable access
	OpcodeLocalGet  Opcode = 0x20
	OpcodeLocalSet  Opcode = 0x21
	OpcodeLocalTee  Opcode = 0x22
	OpcodeGlobalGet Opcode = 0x23
	OpcodeGlobalSet Opcode = 0x24

	// Memory-related operators
	OpcodeI32Load    Opcode = 0x28
	OpcodeI64Load    Opcode = 0x29
	OpcodeF32Load    Opcode = 0x2a
	OpcodeF64Load    Opcode = 0x2b
	OpcodeI32Load8s  Opcode = 0x2c
	OpcodeI32Load8u  Opcode = 0x2d
	OpcodeI32Load16s Opcode = 0x2e
	OpcodeI32Load16u Opcode = 0x2f
	OpcodeI64Load8s  Opcode = 0x30
	OpcodeI64Load8u  Opcode = 0x31
	OpcodeI64Load16s Opcode = 0x32
	OpcodeI64Load16u Opcode = 0x33
	OpcodeI64Load32s Opcode = 0x34
	OpcodeI64Load32u Opcode = 0x35
	OpcodeI32Store   Opcode = 0x36
	OpcodeI64Store   Opcode = 0x37
	OpcodeF32Store   Opcode = 0x38
	OpcodeF64Store   Opcode = 0x39
	OpcodeI32Store8  Opcode = 0x3a
	OpcodeI32Store16 Opcode = 0x3b
	OpcodeI64Store8  Opcode = 0x3c
	OpcodeI64Store16 Opcode = 0x3d
	OpcodeI64Store32 Opcode = 0x3e
	OpcodeMemorySize Opcode = 0x3f
	OpcodeMemoryGrow Opcode = 0x40

	// Constants
	OpcodeI32Const Opcode = 0x41
	OpcodeI64Const Opcode = 0x42
	OpcodeF32Const Opcode = 0x43
	OpcodeF64Const Opcode = 0x44

	// Comparison operators
	OpcodeI32Eqz Opcode = 0x45
	OpcodeI32Eq  Opcode = 0x46
	OpcodeI32Ne  Opcode = 0x47
	OpcodeI32Lts Opcode = 0x48
	OpcodeI32Ltu Opcode = 0x49
	OpcodeI32Gts Opcode = 0x4a
	OpcodeI32Gtu Opcode = 0x4b
	OpcodeI32Les Opcode = 0x4c
	OpcodeI32Leu Opcode = 0x4d
	OpcodeI32Ges Opcode = 0x4e
	OpcodeI32Geu Opcode = 0x4f
	OpcodeI64Eqz Opcode = 0x50
	OpcodeI64Eq  Opcode = 0x51
	OpcodeI64Ne  Opcode = 0x52
	OpcodeI64Lts Opcode = 0x53
	OpcodeI64Ltu Opcode = 0x54
	OpcodeI64Gts Opcode = 0x55
	OpcodeI64Gtu Opcode = 0x56
	OpcodeI64Les Opcode = 0x57
	OpcodeI64Leu Opcode = 0x58
	OpcodeI64Ges Opcode = 0x59
	OpcodeI64Geu Opcode = 0x5a
	OpcodeF32Eq  Opcode = 0x5b
	OpcodeF32Ne  Opcode = 0x5c
	OpcodeF32Lt  Opcode = 0x5d
	OpcodeF32Gt  Opcode = 0x5e
	OpcodeF32Le  Opcode = 0x5f
	OpcodeF32Ge  Opcode = 0x60
	OpcodeF64Eq  Opcode = 0x61
	OpcodeF64Ne  Opcode = 0x62
	OpcodeF64Lt  Opcode = 0x63
	OpcodeF64Gt  Opcode = 0x64
	OpcodeF64Le  Opcode = 0x65
	OpcodeF64Ge  Opcode = 0x66

	// Numeric operators
	OpcodeI32Clz      Opcode = 0x67
	OpcodeI32Ctz      Opcode = 0x68
	OpcodeI32Popcnt   Opcode = 0x69
	OpcodeI32Add      Opcode = 0x6a
	OpcodeI32Sub      Opcode = 0x6b
	OpcodeI32Mul      Opcode = 0x6c
	OpcodeI32Divs     Opcode = 0x6d
	OpcodeI32Divu     Opcode = 0x6e
	OpcodeI32Rems     Opcode = 0x6f
	OpcodeI32Remu     Opcode = 0x70
	OpcodeI32And      Opcode = 0x71
	OpcodeI32Or       Opcode = 0x72
	OpcodeI32Xor      Opcode = 0x73
	OpcodeI32Shl      Opcode = 0x74
	OpcodeI32Shrs     Opcode = 0x75
	OpcodeI32Shru     Opcode = 0x76
	OpcodeI32Rotl     Opcode = 0x77
	OpcodeI32Rotr     Opcode = 0x78
	OpcodeI64Clz      Opcode = 0x79
	OpcodeI64Ctz      Opcode = 0x7a
	OpcodeI64Popcnt   Opcode = 0x7b
	OpcodeI64Add      Opcode = 0x7c
	OpcodeI64Sub      Opcode = 0x7d
	OpcodeI64Mul      Opcode = 0x7e
	OpcodeI64Divs     Opcode = 0x7f
	OpcodeI64Divu     Opcode = 0x80
	OpcodeI64Rems     Opcode = 0x81
	OpcodeI64Remu     Opcode = 0x82
	OpcodeI64And      Opcode = 0x83
	OpcodeI64Or       Opcode = 0x84
	OpcodeI64Xor      Opcode = 0x85
	OpcodeI64Shl      Opcode = 0x86
	OpcodeI64Shrs     Opcode = 0x87
	OpcodeI64Shru     Opcode = 0x88
	OpcodeI64Rotl     Opcode = 0x89
	OpcodeI64Rotr     Opcode = 0x8a
	OpcodeF32Abs      Opcode = 0x8b
	OpcodeF32Neg      Opcode = 0x8c
	OpcodeF32Ceil     Opcode = 0x8d
	OpcodeF32Floor    Opcode = 0x8e
	OpcodeF32Trunc    Opcode = 0x8f
	OpcodeF32Nearest  Opcode = 0x90
	OpcodeF32Sqrt     Opcode = 0x91
	OpcodeF32Add      Opcode = 0x92
	OpcodeF32Sub      Opcode = 0x93
	OpcodeF32Mul      Opcode = 0x94
	OpcodeF32Div      Opcode = 0x95
	OpcodeF32Min      Opcode = 0x96
	OpcodeF32Max      Opcode = 0x97
	OpcodeF32CopySign Opcode = 0x98
	OpcodeF64Abs      Opcode = 0x99
	OpcodeF64Neg      Opcode = 0x9a
	OpcodeF64Ceil     Opcode = 0x9b
	OpcodeF64Floor    Opcode = 0x9c
	OpcodeF64Trunc    Opcode = 0x9d
	OpcodeF64Nearest  Opcode = 0x9e
	OpcodeF64Sqrt     Opcode = 0x9f
	OpcodeF64Add      Opcode = 0xa0
	OpcodeF64Sub      Opcode = 0xa1
	OpcodeF64Mul      Opcode = 0xa2
	OpcodeF64Div      Opcode = 0xa3
	OpcodeF64Min      Opcode = 0xa4
	OpcodeF64Max      Opcode = 0xa5
	OpcodeF64CopySign Opcode = 0xa6

	// Conversions
	OpcodeI32WrapI64     Opcode = 0xa7
	OpcodeI32TruncsF32   Opcode = 0xa8
	OpcodeI32TruncuF32   Opcode = 0xa9
	OpcodeI32TruncsF64   Opcode = 0xaa
	OpcodeI32TruncuF64   Opcode = 0xab
	OpcodeI64ExtendsI32  Opcode = 0xac
	OpcodeI64ExtenduI32  Opcode = 0xad
	OpcodeI64TruncsF32   Opcode = 0xae
	OpcodeI64TruncuF32   Opcode = 0xaf
	OpcodeI64TruncsF64   Opcode = 0xb0
	OpcodeI64TruncuF64   Opcode = 0xb1
	OpcodeF32ConvertsI32 Opcode = 0xb2
	OpcodeF32ConvertuI32 Opcode = 0xb3
	OpcodeF32ConvertsI64 Opcode = 0xb4
	OpcodeF32ConvertuI64 Opcode = 0xb5
	OpcodeF32DemoteF64   Opcode = 0xb6
	OpcodeF64ConvertsI32 Opcode = 0xb7
	OpcodeF64ConvertuI32 Opcode = 0xb8
	OpcodeF64ConvertsI64 Opcode = 0xb9
	OpcodeF64ConvertuI64 Opcode = 0xba
	OpcodeF64PromoteF32  Opcode = 0xbb

	// Reinterpretations
	OpcodeI32ReinterpretF32 Opcode = 0xbc
	OpcodeI64ReinterpretF64 Opcode = 0xbd
	OpcodeF32ReinterpretI32 Opcode = 0xbe
	OpcodeF64ReinterpretI64 Opcode = 0xbf
)

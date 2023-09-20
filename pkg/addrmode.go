package sim6502

type AddressingMode int

const (
	A AddressingMode = iota
	ABS
	ABS_X
	ABS_Y
	IMMED
	IMPL
	IND
	X_IND
	IND_Y
	REL
	ZPG
	ZPG_X
	ZPG_Y
	ZPG_REL   // 65C02 only
	IND_ABS_X // 65C02 only
	ZPG_IND   //65C02 only
)

func (mode AddressingMode) String() string {
	switch mode {
	case A:
		return "A"
	case ABS:
		return "ABS"
	case ABS_X:
		return "ABS_X"
	case ABS_Y:
		return "ABS_Y"
	case IMMED:
		return "IMMEDIATE"
	case IMPL:
		return "IMPLIED"
	case IND:
		return "INDIRECT"
	case X_IND:
		return "X_INDIRECT"
	case IND_Y:
		return "INDIRECT_Y"
	case REL:
		return "RELATIVE"
	case ZPG:
		return "ZPG"
	case ZPG_X:
		return "ZPG_X"
	case ZPG_Y:
		return "ZPG_Y"
	case ZPG_REL:
		return "ZPG_REL"
	case IND_ABS_X:
		return "IND_ASB_X"
	case ZPG_IND:
		return "ZPG_IND"

	default:
		return "???"
	}
}

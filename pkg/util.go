package sim6502

import (
	"fmt"
	"strings"
)

func reverseWord(word uint16) uint16 {
	return (word<<8)&0xff00 | ((word >> 8) & 0xff)
}

func comp2(val uint8) uint8 {
	if val == 0 {
		return 0
	}
	return (val ^ 0xff) + 1
}

func overflows(a, b, carry uint8) bool {

	// The addition of two numbers overflows, if
	// 1. Numbers are of the same sign (additions of opposing signs cannot overflow)
	// 2. The result of the addition yields opposing sign

	// fmt.Printf("OV %d+%d+%d\n", a, b, carry)

	if a&0x80 == b&0x80 {
		// Numbers are of same sign, thus overflow is possible
		r := a + b + carry
		// If the result is not the same sign as the (two) inputs, overflow occurred
		overflowed := a&0x80 != r&0x80
		// fmt.Printf("Result overflow=%v\n", overflowed)
		return overflowed
	}
	// fmt.Printf("Result overflowed=false implicit\n")
	return false
}

func samePage(p1, p2 uint16) bool {
	return (p1 & 0xff00) == (p2 & 0xff00)
}

func FormatInstruction(pc uint16, ins *instruction, b1 uint8, b2 uint8, ma uint16, mc uint8, spre string, cycles uint8, proc *Processor) string {

	var mode string
	var operands string
	var values = []string{
		fmt.Sprintf("A=$%02x", proc.registers.A),
		fmt.Sprintf("X=$%02x", proc.registers.X),
		fmt.Sprintf("Y=$%02x", proc.registers.Y),
		fmt.Sprintf("SP=$%02x", proc.registers.SP.ptr),
		fmt.Sprintf("SH=$%02x", proc.registers.SP.PeekStackHead()),
		fmt.Sprintf("SR=$%02x", proc.registers.SR.Value()),
		fmt.Sprintf("CY=%d", cycles),
		spre,
	}
	switch ins.AddressingMode {
	case A:
		mode = "A"
		operands = ""
	case ABS:
		mode = fmt.Sprintf("$%02x%02x", b2, b1)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case ABS_X:
		mode = fmt.Sprintf("$%02x%02x,X", b2, b1)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case ABS_Y:
		mode = fmt.Sprintf("$%02x%02x,Y", b2, b1)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case IMMED:
		mode = fmt.Sprintf("#$%02x", b1)
		operands = fmt.Sprintf("$%02x", b1)
	case IND:
		mode = fmt.Sprintf("($%02x%02x)", b2, b1)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case X_IND:
		mode = fmt.Sprintf("($%02x,X)", b1)
		operands = fmt.Sprintf("$%02x", b1)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case IND_Y:
		mode = fmt.Sprintf("($%02x),Y", b1)
		operands = fmt.Sprintf("$%02x", b1)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case ZPG:
		mode = fmt.Sprintf("$%02x", b1)
		operands = fmt.Sprintf("$%02x", b1)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case ZPG_X:
		mode = fmt.Sprintf("$%02x,X", b1)
		operands = fmt.Sprintf("$%02x   ", b1)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case ZPG_Y:
		mode = fmt.Sprintf("$%02x,Y", b1)
		operands = fmt.Sprintf("$%02x", b1)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case REL:
		mode = fmt.Sprintf("$%02x", b1)
		operands = fmt.Sprintf("$%02x", b1)
	case IMPL:
		mode = ""
		operands = ""
	case ZPG_REL:
		mode = fmt.Sprintf("$%02x $%02x", b1, b2)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
		values = append(values, fmt.Sprintf("MC=$%02x", mc))
	case IND_ABS_X:
		mode = fmt.Sprintf("($%02x%02x,X)", b2, b1)
		operands = fmt.Sprintf("$%02x $%02x", b1, b2)
		values = append(values, fmt.Sprintf("MA=$%04x", ma))
	case ZPG_IND:
		mode = fmt.Sprintf("($%02x)", b1)
		operands = fmt.Sprintf("$%02x", b1)
		values = append(values, fmt.Sprintf("MA=%04x", ma))
	default:
		mode = fmt.Sprintf("?%d", ins.AddressingMode)
		operands = ""
	}

	return fmt.Sprintf("$%04x: $%02x %-7s : %-4s %-10s  [%s]", pc, ins.OpCode, operands, ins.Impl.Mnemonic(), mode, strings.Join(values, " "))
}

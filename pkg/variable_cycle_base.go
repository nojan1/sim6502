package sim6502

// variableInstructionBranch is as a base by branch instructions to
// tweak cycles depending on branch target
type variableInstructionBranch struct{}

func (v *variableInstructionBranch) TweakCycles(m AddressingMode, instructionPC uint16, address uint16, xPageBoundary bool) uint8 {
	// Branches are always address mode relative, so we don't need to check the mode

	// If we're branching accross a page, we need to increment by 2, otherwise 1
	if samePage(instructionPC, address) {
		return 1
	}
	return 2
}

// standardCrossPage is implemented by instructions with ABS_X, ABS_Y and IND_Y
// modes where a page is crossed, to increment cycles by one
type standardCrossPage struct{}

func (v *standardCrossPage) TweakCycles(m AddressingMode, instructionPC uint16, address uint16, xPageBoundary bool) uint8 {
	if xPageBoundary {
		return 1
	}
	return 0
}

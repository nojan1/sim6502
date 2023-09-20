package sim6502

// Adds 85C02 specific instructions

// loadExtendedInstructions will load the extended instructions into an instruction map
func load85C02instructions(instructions []*instruction) {

	// New instructions
	loadAdditionalInstruction(instructions, 0x80, &bra{}, REL, 3)

	loadAdditionalInstruction(instructions, 0xda, &phx{}, IMPL, 3)

	loadAdditionalInstruction(instructions, 0xfa, &plx{}, IMPL, 4)

	loadAdditionalInstruction(instructions, 0x5a, &phy{}, IMPL, 3)

	loadAdditionalInstruction(instructions, 0x7a, &ply{}, IMPL, 4)

	loadAdditionalInstruction(instructions, 0x64, &stz{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0x74, &stz{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x9c, &stz{}, ABS, 4)
	loadAdditionalInstruction(instructions, 0x9e, &stz{}, ABS_X, 5)

	loadAdditionalInstruction(instructions, 0x14, &trb{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x1c, &trb{}, ABS, 6)

	loadAdditionalInstruction(instructions, 0x04, &tsb{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x0c, &tsb{}, ABS, 6)

	loadAdditionalInstruction(instructions, 0x0f, &bbr0{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x1f, &bbr1{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x2f, &bbr2{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x3f, &bbr3{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x4f, &bbr4{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x5f, &bbr5{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x6f, &bbr6{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x7f, &bbr7{}, ZPG_REL, 5)

	loadAdditionalInstruction(instructions, 0x8f, &bbs0{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0x9f, &bbs1{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xaf, &bbs2{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xbf, &bbs3{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xcf, &bbs4{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xdf, &bbs5{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xef, &bbs6{}, ZPG_REL, 5)
	loadAdditionalInstruction(instructions, 0xff, &bbs7{}, ZPG_REL, 5)

	loadAdditionalInstruction(instructions, 0x07, &rmb0{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x17, &rmb1{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x27, &rmb2{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x37, &rmb3{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x47, &rmb4{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x57, &rmb5{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x67, &rmb6{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x77, &rmb7{}, ZPG, 5)

	loadAdditionalInstruction(instructions, 0x87, &smb0{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x97, &smb1{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xa7, &smb2{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xb7, &smb3{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xc7, &smb4{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xd7, &smb5{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xe7, &smb6{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xf7, &smb7{}, ZPG, 5)

	loadAdditionalInstruction(instructions, 0xdb, &stp{}, IMPL, 3)

	loadAdditionalInstruction(instructions, 0xcb, &wai{}, IMPL, 3)

	// Additional implementation modes for standard 6502 instructions
	loadAdditionalInstruction(instructions, 0x72, &adc{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0x32, &and{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0xd2, &cmp{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0x52, &eor{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0xb2, &lda{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0x12, &ora{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0xf2, &sbc{}, ZPG_IND, 3)
	loadAdditionalInstruction(instructions, 0x92, &sta{}, ZPG_IND, 3)

	loadAdditionalInstruction(instructions, 0x89, &bit{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0x34, &bit{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x3c, &bit{}, ABS_X, 4)

	loadAdditionalInstruction(instructions, 0x3a, &dec{}, A, 2)
	loadAdditionalInstruction(instructions, 0x1a, &inc{}, A, 2)

	loadAdditionalInstruction(instructions, 0x7c, &jmp{}, IND_ABS_X, 6)

	// Everything else is a NOP
	// The NOP instruction will consume additional bytes as needed
	// for those NOPs that are multi-byte
	for i, ins := range instructions {
		if ins == nil {
			loadAdditionalInstruction(instructions, uint8(i), &nop{}, IMPL, 1)
		}
	}
}

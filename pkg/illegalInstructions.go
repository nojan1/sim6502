package sim6502

// loadIllegalInstructions will load the illegal instructions into an instruction map
func loadIllegalInstructions(instructions []*instruction) {

	// JAMs
	loadAdditionalInstruction(instructions, 0x02, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x12, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x22, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x32, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x42, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x52, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x62, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x72, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x92, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0xb2, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0xd2, &jam{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0xf2, &jam{}, IMPL, 2)

	// NOPs
	loadAdditionalInstruction(instructions, 0x1a, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x3a, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x5a, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x7a, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0xda, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0xfa, &nop{}, IMPL, 2)
	loadAdditionalInstruction(instructions, 0x80, &nop{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0x82, &nop{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0x89, &nop{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0xc2, &nop{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0xe2, &nop{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0x04, &nop{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0x44, &nop{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0x64, &nop{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0x14, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x34, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x54, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x74, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0xd4, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0xf4, &nop{}, ZPG_X, 4)
	loadAdditionalInstruction(instructions, 0x0c, &nop{}, ABS, 4)
	loadAdditionalInstruction(instructions, 0x1c, &nop{}, ABS_X, 4)
	loadAdditionalInstruction(instructions, 0x3c, &nop{}, ABS_X, 4)
	loadAdditionalInstruction(instructions, 0x5c, &nop{}, ABS_X, 4)
	loadAdditionalInstruction(instructions, 0x7c, &nop{}, ABS_X, 4)
	loadAdditionalInstruction(instructions, 0xdc, &nop{}, ABS_X, 4)
	loadAdditionalInstruction(instructions, 0xfc, &nop{}, ABS_X, 4)

	// ALR
	loadAdditionalInstruction(instructions, 0x4b, &alr{}, IMMED, 2)

	// ANC
	loadAdditionalInstruction(instructions, 0x0b, &anc{}, IMMED, 2)
	loadAdditionalInstruction(instructions, 0x2b, &anc{}, IMMED, 2)

	// ANE
	loadAdditionalInstruction(instructions, 0x8b, &ane{}, IMMED, 2)

	// ARR
	loadAdditionalInstruction(instructions, 0x6b, &arr{}, IMMED, 2)

	// DCP
	loadAdditionalInstruction(instructions, 0xc7, &dcp{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xd7, &dcp{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0xcf, &dcp{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0xdf, &dcp{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0xdb, &dcp{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0xc3, &dcp{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0xd3, &dcp{}, IND_Y, 8)

	// ISC
	loadAdditionalInstruction(instructions, 0xe7, &isc{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0xf7, &isc{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0xef, &isc{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0xff, &isc{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0xfb, &isc{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0xe3, &isc{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0xf3, &isc{}, IND_Y, 8)

	// LAS
	loadAdditionalInstruction(instructions, 0xbb, &las{}, ABS_Y, 4)

	// LAX
	loadAdditionalInstruction(instructions, 0xa7, &lax{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0xb7, &lax{}, ZPG_Y, 4)
	loadAdditionalInstruction(instructions, 0xaf, &lax{}, ABS, 4)
	loadAdditionalInstruction(instructions, 0xbf, &lax{}, ABS_Y, 4)
	loadAdditionalInstruction(instructions, 0xa3, &lax{}, X_IND, 6)
	loadAdditionalInstruction(instructions, 0xb3, &lax{}, IND_Y, 5)

	// LXA
	loadAdditionalInstruction(instructions, 0xab, &lxa{}, IMMED, 2)

	// RLA
	loadAdditionalInstruction(instructions, 0x27, &rla{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x37, &rla{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0x2f, &rla{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0x3f, &rla{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0x3b, &rla{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0x23, &rla{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0x33, &rla{}, IND_Y, 8)

	// RRA
	loadAdditionalInstruction(instructions, 0x67, &rra{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x77, &rra{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0x6f, &rra{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0x7f, &rra{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0x7b, &rra{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0x63, &rra{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0x73, &rra{}, IND_Y, 8)

	// SAX
	loadAdditionalInstruction(instructions, 0x87, &sax{}, ZPG, 3)
	loadAdditionalInstruction(instructions, 0x97, &sax{}, ZPG_Y, 4)
	loadAdditionalInstruction(instructions, 0x8f, &sax{}, ABS, 4)
	loadAdditionalInstruction(instructions, 0x83, &sax{}, X_IND, 6)

	// SBX
	loadAdditionalInstruction(instructions, 0xcb, &sbx{}, IMMED, 2)

	// SHA
	loadAdditionalInstruction(instructions, 0x9f, &sha{}, ABS_Y, 5)
	loadAdditionalInstruction(instructions, 0x93, &sha{}, IND_Y, 6)

	// SHX
	loadAdditionalInstruction(instructions, 0x9e, &shx{}, ABS_Y, 5)

	// SHY
	loadAdditionalInstruction(instructions, 0x9c, &shy{}, ABS_X, 5)

	// SLO
	loadAdditionalInstruction(instructions, 0x07, &slo{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x17, &slo{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0x0f, &slo{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0x1f, &slo{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0x1b, &slo{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0x03, &slo{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0x13, &slo{}, IND_Y, 8)

	// SRE
	loadAdditionalInstruction(instructions, 0x47, &sre{}, ZPG, 5)
	loadAdditionalInstruction(instructions, 0x57, &sre{}, ZPG_X, 6)
	loadAdditionalInstruction(instructions, 0x4f, &sre{}, ABS, 6)
	loadAdditionalInstruction(instructions, 0x5f, &sre{}, ABS_X, 7)
	loadAdditionalInstruction(instructions, 0x5b, &sre{}, ABS_Y, 7)
	loadAdditionalInstruction(instructions, 0x43, &sre{}, X_IND, 8)
	loadAdditionalInstruction(instructions, 0x53, &sre{}, IND_Y, 8)

	// TAS
	loadAdditionalInstruction(instructions, 0x9b, &tas{}, ABS_Y, 5)

	// USB (USBC)
	loadAdditionalInstruction(instructions, 0xeb, &usb{}, IMMED, 2)

}

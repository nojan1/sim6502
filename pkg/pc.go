package sim6502

type ProgramCounter struct {
	ptr       uint16
	processor *Processor
}

func (pc *ProgramCounter) Init(p *Processor) {
	pc.processor = p
	pc.ptr = uint16(p.memory.Read(uint16(VectorReset))) | (uint16(p.memory.Read(uint16(VectorReset+1))) << 8)
}

func (pc *ProgramCounter) Next() uint8 {
	next := pc.processor.memory.Read(pc.ptr)
	pc.ptr++
	return next
}

func (pc *ProgramCounter) Current() uint16 {
	return pc.ptr
}

func (pc *ProgramCounter) Set(val uint16) {
	pc.ptr = val
}

package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs int32
	matchOffsets []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for index := int32(0); index < self.npairs * 2; index += 2 {
		if self.matchOffsets[index] == key {
			offset := self.matchOffsets[index]
			base.Branch(frame, int(offset))
			return 
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}

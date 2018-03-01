package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type GOTO struct {base.BranchInstruction}

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}
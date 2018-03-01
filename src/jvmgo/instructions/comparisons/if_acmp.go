package comparisons

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IF_ACMPEQ struct {base.BranchInstruction}
type IF_ACMPNE struct {base.BranchInstruction}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame)  {
	val1 := frame.OperandStack().PopRef()
	val2 := frame.OperandStack().PopRef()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame)  {
	val1 := frame.OperandStack().PopRef()
	val2 := frame.OperandStack().PopRef()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
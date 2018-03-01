package conversions

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type F2I struct {base.NoOperandsInstruction}
type F2D struct {base.NoOperandsInstruction}
type F2L struct {base.NoOperandsInstruction}

func (self *F2I) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int32(d)
	stack.PushInt(i)
}

func (self *F2D) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float64(d)
	stack.PushDouble(i)
}

func (self *F2L) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int64(d)
	stack.PushLong(i)
}

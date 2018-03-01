package conversions

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type L2F struct {base.NoOperandsInstruction}
type L2D struct {base.NoOperandsInstruction}
type L2I struct {base.NoOperandsInstruction}

func (self *L2F) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float32(d)
	stack.PushFloat(i)
}

func (self *L2D) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float64(d)
	stack.PushDouble(i)
}

func (self *L2I) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := int32(d)
	stack.PushInt(i)
}

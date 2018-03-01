package math

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DNEG struct {base.NoOperandsInstruction}
type FNEG struct {base.NoOperandsInstruction}
type INEG struct {base.NoOperandsInstruction}
type LNEG struct {base.NoOperandsInstruction}

func (self *INEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (self *LNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

func (self *DNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

func (self *FNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

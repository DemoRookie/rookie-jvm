package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// 空操作
type NoOperandsInstruction struct {

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	
}

// 跳转
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}

// 局部变量存取
type Index8Instruction struct {
	Index uint // 局部变量索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
}

// 常量池存取
type Index16Instruction struct {
	Index uint // 常量索引
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadUint16())
}

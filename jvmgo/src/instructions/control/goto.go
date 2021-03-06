package control

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type GOTO struct { base.BranchInstruction}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
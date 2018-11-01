package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (c *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return c.constantValueIndex
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

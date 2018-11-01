package classfile

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (e *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {

	self.exceptionIndexTable = reader.readUint16s()
}

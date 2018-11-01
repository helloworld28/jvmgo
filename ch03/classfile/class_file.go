package classfile

import "fmt"

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (c *ClassFile) Attributes() []AttributeInfo {
	return c.attributes
}

func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {

}

func (self *ClassFile) MinorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}
func (self *ClassFile) InterfaceNames() []string {
	interFaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interFaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interFaceNames
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

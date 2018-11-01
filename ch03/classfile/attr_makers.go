package classfile

type DeprecatedAttribute struct {
	MarkAttribute
}
type SyntheticAttribute struct {
	MarkAttribute
}
type MarkAttribute struct {
}

func (self *MarkAttribute) readInfo(reader *ClassReader) {

}

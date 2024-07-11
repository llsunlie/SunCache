package member

type ByteView struct {
	value []byte
}

func (b ByteView) Len() (length int) {
	return len(b.value)
}

func (b ByteView) Copy() (byteView []byte) {
	return cloneBytes(b.value)
}

func (b ByteView) String() (str string) {
	return string(b.value)
}

func cloneBytes(a []byte) (b []byte) {
	b = make([]byte, len(a))
	copy(b, a)
	return
}

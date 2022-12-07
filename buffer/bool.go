package buffer

func (b *WriteBuffer) WriteBool(v bool) {
	if v {
		b.buf.WriteByte(1)
	} else {
		b.buf.WriteByte(0)
	}
}

func (b *ReadBuffer) ReadBool() (bool, error) {
	v := b.buf[0]
	b.buf = b.buf[1:]
	return !(v == 0), nil
}

func (b *ReadBuffer) SkipBool() {
	b.buf = b.buf[1:]
}

func (b *ReadBuffer) ChangeBool(v bool) {
	if v {
		b.buf[0] = 1
	} else {
		b.buf[0] = 0
	}
}

package buffer

import "encoding/binary"

func (b *WriteBuffer) WriteString(s string) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], uint32(len(s)))
	b.buf.Write(l[:])
	b.buf.WriteString(s)
}

func (b *ReadBuffer) ReadString() (string, error) {
	l := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4:]
	v := string(b.buf[:l])
	b.buf = b.buf[l:]
	return v, nil
}

func (b *ReadBuffer) SkipString() {
	l := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4+int(l):]
}

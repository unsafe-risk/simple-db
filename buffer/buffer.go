package buffer

import (
	"bytes"
	"encoding/binary"
)

type WriteBuffer struct {
	buf *bytes.Buffer
}

func NewWriteBuffer() *WriteBuffer {
	return &WriteBuffer{
		buf: bytes.NewBuffer(nil),
	}
}

func (b *WriteBuffer) Write(v []byte) error {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], uint32(len(v)))
	b.buf.Write(l[:])
	b.buf.Write(v)
	return nil
}

func (b *WriteBuffer) Bytes() []byte {
	return b.buf.Bytes()
}

type ReadBuffer struct {
	buf []byte
}

func NewReadBuffer(v []byte) *ReadBuffer {
	return &ReadBuffer{
		buf: v,
	}
}

func (b *ReadBuffer) Read() ([]byte, error) {
	l := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4:]
	v := make([]byte, l)
	copy(v, b.buf[:l])
	b.buf = b.buf[l:]
	return v, nil
}

func (b *ReadBuffer) SkipBytes() {
	l := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4+int(l):]
}

type ModifyBuffer struct {
	buf    []byte
	result *bytes.Buffer
}

func NewModifyBuffer(v []byte) *ModifyBuffer {
	return &ModifyBuffer{
		buf:    v,
		result: bytes.NewBuffer(make([]byte, 0, len(v))),
	}
}

func (b *ModifyBuffer) SkipBytes() {
	b.result.Write(b.buf[:4])
	l := binary.BigEndian.Uint32(b.buf[:4])
	b.result.Write(b.buf[4 : 4+l])
	b.buf = b.buf[4+int(l):]
}

func (b *ModifyBuffer) Modify(v []byte) {
	b.result.Write(b.buf[:4])
	u := binary.BigEndian.Uint32(b.buf[:4])
	b.result.Write(b.buf[4 : 4+u])
	b.buf = b.buf[4+int(u):]
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], uint32(len(v)))
	b.result.Write(l[:])
	b.result.Write(v)
}

func (b *ModifyBuffer) Result() []byte {
	return b.result.Bytes()
}

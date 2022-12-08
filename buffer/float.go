package buffer

import (
	"encoding/binary"
	"math"
)

func (b *WriteBuffer) WriteFloat32(v float32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], math.Float32bits(v))
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteFloat64(v float64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], math.Float64bits(v))
	b.buf.Write(l[:])
}

func (b *ReadBuffer) ReadFloat32() (float32, error) {
	v := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4:]
	return math.Float32frombits(v), nil
}

func (b *ReadBuffer) ReadFloat64() (float64, error) {
	v := binary.BigEndian.Uint64(b.buf[:8])
	b.buf = b.buf[8:]
	return math.Float64frombits(v), nil
}

func (b *ReadBuffer) SkipFloat32() {
	b.buf = b.buf[4:]
}

func (b *ReadBuffer) SkipFloat64() {
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) SkipFloat32() {
	b.result.Write(b.buf[:4])
	b.buf = b.buf[4:]
}

func (b *ModifyBuffer) SkipFloat64() {
	b.result.Write(b.buf[:8])
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) ModifyFloat32(f float32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], math.Float32bits(f))
	b.result.Write(l[:])
}

func (b *ModifyBuffer) ModifyFloat64(f float64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], math.Float64bits(f))
	b.result.Write(l[:])
}

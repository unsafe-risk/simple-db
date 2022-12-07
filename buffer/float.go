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

func (b *ReadBuffer) ChangeFloat32(v float32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], math.Float32bits(v))
	b.buf[0] = l[0]
	b.buf[1] = l[1]
	b.buf[2] = l[2]
	b.buf[3] = l[3]
}

func (b *ReadBuffer) ChangeFloat64(v float64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], math.Float64bits(v))
	b.buf[0] = l[0]
	b.buf[1] = l[1]
	b.buf[2] = l[2]
	b.buf[3] = l[3]
	b.buf[4] = l[4]
	b.buf[5] = l[5]
	b.buf[6] = l[6]
	b.buf[7] = l[7]
}

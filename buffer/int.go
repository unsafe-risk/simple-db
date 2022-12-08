package buffer

import (
	"encoding/binary"
)

func (b *WriteBuffer) WriteInt8(v int8) {
	b.buf.WriteByte(byte(v))
}

func (b *WriteBuffer) WriteInt16(v int16) {
	l := [2]byte{}
	binary.BigEndian.PutUint16(l[:], uint16(v))
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteInt32(v int32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], uint32(v))
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteInt64(v int64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], uint64(v))
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteUint8(v uint8) {
	b.buf.WriteByte(v)
}

func (b *WriteBuffer) WriteUint16(v uint16) {
	l := [2]byte{}
	binary.BigEndian.PutUint16(l[:], v)
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteUint32(v uint32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], v)
	b.buf.Write(l[:])
}

func (b *WriteBuffer) WriteUint64(v uint64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], v)
	b.buf.Write(l[:])
}

func (b *ReadBuffer) ReadInt8() (int8, error) {
	v := b.buf[0]
	b.buf = b.buf[1:]
	return int8(v), nil
}

func (b *ReadBuffer) ReadInt16() (int16, error) {
	v := binary.BigEndian.Uint16(b.buf[:2])
	b.buf = b.buf[2:]
	return int16(v), nil
}

func (b *ReadBuffer) ReadInt32() (int32, error) {
	v := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4:]
	return int32(v), nil
}

func (b *ReadBuffer) ReadInt64() (int64, error) {
	v := binary.BigEndian.Uint64(b.buf[:8])
	b.buf = b.buf[8:]
	return int64(v), nil
}

func (b *ReadBuffer) ReadUint8() (uint8, error) {
	v := b.buf[0]
	b.buf = b.buf[1:]
	return v, nil
}

func (b *ReadBuffer) ReadUint16() (uint16, error) {
	v := binary.BigEndian.Uint16(b.buf[:2])
	b.buf = b.buf[2:]
	return v, nil
}

func (b *ReadBuffer) ReadUint32() (uint32, error) {
	v := binary.BigEndian.Uint32(b.buf[:4])
	b.buf = b.buf[4:]
	return v, nil
}

func (b *ReadBuffer) ReadUint64() (uint64, error) {
	v := binary.BigEndian.Uint64(b.buf[:8])
	b.buf = b.buf[8:]
	return v, nil
}

func (b *ReadBuffer) SkipInt8() {
	b.buf = b.buf[1:]
}

func (b *ReadBuffer) SkipInt16() {
	b.buf = b.buf[2:]
}

func (b *ReadBuffer) SkipInt32() {
	b.buf = b.buf[4:]
}

func (b *ReadBuffer) SkipInt64() {
	b.buf = b.buf[8:]
}

func (b *ReadBuffer) SkipUint8() {
	b.buf = b.buf[1:]
}

func (b *ReadBuffer) SkipUint16() {
	b.buf = b.buf[2:]
}

func (b *ReadBuffer) SkipUint32() {
	b.buf = b.buf[4:]
}

func (b *ReadBuffer) SkipUint64() {
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) SkipInt8() {
	b.result.WriteByte(b.buf[0])
	b.buf = b.buf[1:]
}

func (b *ModifyBuffer) SkipInt16() {
	b.result.Write(b.buf[:2])
	b.buf = b.buf[2:]
}

func (b *ModifyBuffer) SkipInt32() {
	b.result.Write(b.buf[:4])
	b.buf = b.buf[4:]
}

func (b *ModifyBuffer) SkipInt64() {
	b.result.Write(b.buf[:8])
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) SkipUint8() {
	b.result.WriteByte(b.buf[0])
	b.buf = b.buf[1:]
}

func (b *ModifyBuffer) SkipUint16() {
	b.result.Write(b.buf[:2])
	b.buf = b.buf[2:]
}

func (b *ModifyBuffer) SkipUint32() {
	b.result.Write(b.buf[:4])
	b.buf = b.buf[4:]
}

func (b *ModifyBuffer) SkipUint64() {
	b.result.Write(b.buf[:8])
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) ModifyInt8(i int8) {
	b.result.WriteByte(byte(i))
	b.buf = b.buf[1:]
}

func (b *ModifyBuffer) ModifyInt16(i int16) {
	l := [2]byte{}
	binary.BigEndian.PutUint16(l[:], uint16(i))
	b.result.Write(l[:])
	b.buf = b.buf[2:]
}

func (b *ModifyBuffer) ModifyInt32(i int32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], uint32(i))
	b.result.Write(l[:])
	b.buf = b.buf[4:]
}

func (b *ModifyBuffer) ModifyInt64(i int64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], uint64(i))
	b.result.Write(l[:])
	b.buf = b.buf[8:]
}

func (b *ModifyBuffer) ModifyUint8(i uint8) {
	b.result.WriteByte(i)
	b.buf = b.buf[1:]
}

func (b *ModifyBuffer) ModifyUint16(i uint16) {
	l := [2]byte{}
	binary.BigEndian.PutUint16(l[:], i)
	b.result.Write(l[:])
	b.buf = b.buf[2:]
}

func (b *ModifyBuffer) ModifyUint32(i uint32) {
	l := [4]byte{}
	binary.BigEndian.PutUint32(l[:], i)
	b.result.Write(l[:])
	b.buf = b.buf[4:]
}

func (b *ModifyBuffer) ModifyUint64(i uint64) {
	l := [8]byte{}
	binary.BigEndian.PutUint64(l[:], i)
	b.result.Write(l[:])
	b.buf = b.buf[8:]
}

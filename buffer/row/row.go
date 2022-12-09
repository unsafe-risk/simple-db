package row

import (
	"errors"
	"github.com/unsafe-risk/simple-db/buffer"
	"github.com/unsafe-risk/simple-db/buffer/column"
)

type Row struct {
	typeList []int
	buf      []byte
}

func New(typeList ...int) *Row {
	return &Row{
		typeList: typeList,
	}
}

func NewBlank(typeList ...int) (*Row, error) {
	r := New(typeList...)
	wr := buffer.NewWriteBuffer()
	for _, t := range typeList {
		switch t {
		case column.Bool:
			wr.WriteBool(false)
		case column.Int64:
			wr.WriteInt64(0)
		case column.Int32:
			wr.WriteInt32(0)
		case column.Int16:
			wr.WriteInt16(0)
		case column.Int8:
			wr.WriteInt8(0)
		case column.Uint64:
			wr.WriteUint64(0)
		case column.Uint32:
			wr.WriteUint32(0)
		case column.Uint16:
			wr.WriteUint16(0)
		case column.Uint8:
			wr.WriteUint8(0)
		case column.Float64:
			wr.WriteFloat64(0)
		case column.Float32:
			wr.WriteFloat32(0)
		case column.String:
			wr.WriteString("")
		case column.Bytes:
			wr.Write([]byte{})
		default:
			return nil, errors.New("unknown column type")
		}
	}
	r.buf = wr.Bytes()
	return r, nil
}

func (r *Row) SetBytes(v []byte) error {
	r.buf = v
	return nil
}

func Get[T column.Column](r *Row, i int) (rs T, err error) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("index out of range")
			}
		}()
		t := buffer.NewReadBuffer(r.buf)
		for j := 0; j < i; j++ {
			switch r.typeList[j] {
			case column.Bool:
				t.SkipBool()
			case column.Int64:
				t.SkipInt64()
			case column.Int32:
				t.SkipInt32()
			case column.Int16:
				t.SkipInt16()
			case column.Int8:
				t.SkipInt8()
			case column.Uint64:
				t.SkipUint64()
			case column.Uint32:
				t.SkipUint32()
			case column.Uint16:
				t.SkipUint16()
			case column.Uint8:
				t.SkipUint8()
			case column.Float64:
				t.SkipFloat64()
			case column.Float32:
				t.SkipFloat32()
			case column.String:
				t.SkipString()
			case column.Bytes:
				t.SkipBytes()
			default:
				err = errors.New("unknown type")
				return
			}
		}
		switch r.typeList[i] {
		case column.Bool:
			v, e := t.ReadBool()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int64:
			v, e := t.ReadInt64()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int32:
			v, e := t.ReadInt32()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int16:
			v, e := t.ReadInt16()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int8:
			v, e := t.ReadInt8()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint64:
			v, e := t.ReadUint64()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint32:
			v, e := t.ReadUint32()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint16:
			v, e := t.ReadUint16()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint8:
			v, e := t.ReadUint8()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Float64:
			v, e := t.ReadFloat64()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Float32:
			v, e := t.ReadFloat32()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.String:
			v, e := t.ReadString()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Bytes:
			v, e := t.Read()
			if e != nil {
				err = e
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		default:
			err = errors.New("unknown type")
			return
		}
	}()
	return
}

func Modify[T column.Column](r *Row, i int, v T) (err error) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("index out of range")
			}
		}()
		t := buffer.NewModifyBuffer(r.buf)
		for j := 0; j < i; j++ {
			switch r.typeList[j] {
			case column.Bool:
				t.SkipBool()
			case column.Int64:
				t.SkipInt64()
			case column.Int32:
				t.SkipInt32()
			case column.Int16:
				t.SkipInt16()
			case column.Int8:
				t.SkipInt8()
			case column.Uint64:
				t.SkipUint64()
			case column.Uint32:
				t.SkipUint32()
			case column.Uint16:
				t.SkipUint16()
			case column.Uint8:
				t.SkipUint8()
			case column.Float64:
				t.SkipFloat64()
			case column.Float32:
				t.SkipFloat32()
			case column.String:
				t.SkipString()
			case column.Bytes:
				t.SkipBytes()
			default:
				err = errors.New("unknown type")
				return
			}
		}
		a := any(v)
		switch r.typeList[i] {
		case column.Bool:
			p, ok := a.(bool)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyBool(p)
		case column.Int64:
			p, ok := a.(int64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyInt64(p)
		case column.Int32:
			p, ok := a.(int32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyInt32(p)
		case column.Int16:
			p, ok := a.(int16)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyInt16(p)
		case column.Int8:
			p, ok := a.(int8)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyInt8(p)
		case column.Uint64:
			p, ok := a.(uint64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyUint64(p)
		case column.Uint32:
			p, ok := a.(uint32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyUint32(p)
		case column.Uint16:
			p, ok := a.(uint16)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyUint16(p)
		case column.Uint8:
			p, ok := a.(uint8)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyUint8(p)
		case column.Float64:
			p, ok := a.(float64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyFloat64(p)
		case column.Float32:
			p, ok := a.(float32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyFloat32(p)
		case column.String:
			p, ok := a.(string)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ModifyString(p)
		case column.Bytes:
			p, ok := a.([]byte)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.Modify(p)
		default:
			err = errors.New("unknown type")
			return
		}
		for j := i + 1; j < len(r.typeList); j++ {
			switch r.typeList[j] {
			case column.Bool:
				t.SkipBool()
			case column.Int64:
				t.SkipInt64()
			case column.Int32:
				t.SkipInt32()
			case column.Int16:
				t.SkipInt16()
			case column.Int8:
				t.SkipInt8()
			case column.Uint64:
				t.SkipUint64()
			case column.Uint32:
				t.SkipUint32()
			case column.Uint16:
				t.SkipUint16()
			case column.Uint8:
				t.SkipUint8()
			case column.Float64:
				t.SkipFloat64()
			case column.Float32:
				t.SkipFloat32()
			case column.String:
				t.SkipString()
			case column.Bytes:
				t.SkipBytes()
			default:
				err = errors.New("unknown type")
				return
			}
		}
		r.buf = t.Result()
	}()
	return
}

func (r *Row) GetBytes() []byte {
	p := make([]byte, len(r.buf))
	copy(p, r.buf)
	return p
}

func (r *Row) EqualColumnTypes(t ...int) bool {
	if len(t) != len(r.typeList) {
		return false
	}
	for i := range t {
		if t[i] != r.typeList[i] {
			return false
		}
	}
	return true
}

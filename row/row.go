package row

import (
	"errors"

	"github.com/unsafe-risk/simple-db/buffer"
	"github.com/unsafe-risk/simple-db/row/column"
)

type Row struct {
	typeList []int
	buf      []byte
}

func NewRow(typeList []int) *Row {
	return &Row{
		typeList: typeList,
	}
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
			v, err := t.ReadBool()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int64:
			v, err := t.ReadInt64()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int32:
			v, err := t.ReadInt32()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int16:
			v, err := t.ReadInt16()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Int8:
			v, err := t.ReadInt8()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint64:
			v, err := t.ReadUint64()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint32:
			v, err := t.ReadUint32()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint16:
			v, err := t.ReadUint16()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Uint8:
			v, err := t.ReadUint8()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Float64:
			v, err := t.ReadFloat64()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Float32:
			v, err := t.ReadFloat32()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.String:
			v, err := t.ReadString()
			if err != nil {
				return
			}
			p, ok := any(v).(T)
			if !ok {
				err = errors.New("type mismatch")
				return
			}
			rs = p
		case column.Bytes:
			v, err := t.Read()
			if err != nil {
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

func Change[T column.Changable](r *Row, i int, v T) (err error) {
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
		a := any(v)
		switch r.typeList[i] {
		case column.Bool:
			p, ok := a.(bool)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeBool(p)
		case column.Int64:
			p, ok := a.(int64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeInt64(p)
		case column.Int32:
			p, ok := a.(int32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeInt32(p)
		case column.Int16:
			p, ok := a.(int16)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeInt16(p)
		case column.Int8:
			p, ok := a.(int8)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeInt8(p)
		case column.Uint64:
			p, ok := a.(uint64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeUint64(p)
		case column.Uint32:
			p, ok := a.(uint32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeUint32(p)
		case column.Uint16:
			p, ok := a.(uint16)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeUint16(p)
		case column.Uint8:
			p, ok := a.(uint8)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeUint8(p)
		case column.Float64:
			p, ok := a.(float64)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeFloat64(p)
		case column.Float32:
			p, ok := a.(float32)
			if !ok {
				err = errors.New("type error")
				return
			}
			t.ChangeFloat32(p)
		default:
			err = errors.New("unknown type")
			return
		}
	}()
	return
}

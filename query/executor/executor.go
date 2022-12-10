package executor

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
	"github.com/unsafe-risk/simple-db/buffer/column"
	"github.com/unsafe-risk/simple-db/buffer/row"
	"github.com/unsafe-risk/simple-db/buffer/row/table"
	"github.com/unsafe-risk/simple-db/query"
	"strconv"
)

var (
	okResponse   = []byte("OK")
	failResponse = []byte("FAIL")
)

type Executor struct {
	tb *table.Table
}

func New() (*Executor, error) {
	exe := &Executor{}
	return exe, nil
}

func (e *Executor) SetTable(tableName string, columnNames []string, columnTypes []int) error {
	tb, err := table.New("database/" + tableName)
	if err != nil {
		return err
	}

	if !tb.SetColumns(columnNames, columnTypes) {
		return errors.New("failed to set columns")
	}

	e.tb = tb

	return nil
}

func (e *Executor) Close() error {
	return e.tb.Close()
}

func (e *Executor) Run(data []byte) ([]byte, error) {
	if e.tb == nil {
		return nil, errors.New("table not set")
	}

	q, err := query.Parse(data)
	if err != nil {
		return nil, err
	}

	switch q.Type {
	case query.SetQuery:
		r, err := e.tb.GetRow(q.Key)
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				r, err = e.tb.GetBlankRow()
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
		switch e.tb.GetColumnType(q.Column) {
		case column.Bool:
			b, err := strconv.ParseBool(string(q.Value))
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), b); err != nil {
				return nil, err
			}
		case column.Bytes:
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), q.Value); err != nil {
				return nil, err
			}
		case column.String:
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), string(q.Value)); err != nil {
				return nil, err
			}
		case column.Int8:
			v, err := strconv.ParseInt(string(q.Value), 10, 8)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), int8(v)); err != nil {
				return nil, err
			}
		case column.Int16:
			v, err := strconv.ParseInt(string(q.Value), 10, 16)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), int16(v)); err != nil {
				return nil, err
			}
		case column.Int32:
			v, err := strconv.ParseInt(string(q.Value), 10, 32)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), int32(v)); err != nil {
				return nil, err
			}
		case column.Int64:
			v, err := strconv.ParseInt(string(q.Value), 10, 64)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), v); err != nil {
				return nil, err
			}
		case column.Uint8:
			v, err := strconv.ParseUint(string(q.Value), 10, 8)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), uint8(v)); err != nil {
				return nil, err
			}
		case column.Uint16:
			v, err := strconv.ParseUint(string(q.Value), 10, 16)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), uint16(v)); err != nil {
				return nil, err
			}
		case column.Uint32:
			v, err := strconv.ParseUint(string(q.Value), 10, 32)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), uint32(v)); err != nil {
				return nil, err
			}
		case column.Uint64:
			v, err := strconv.ParseUint(string(q.Value), 10, 64)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), v); err != nil {
				return nil, err
			}
		case column.Float32:
			v, err := strconv.ParseFloat(string(q.Value), 32)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), float32(v)); err != nil {
				return nil, err
			}
		case column.Float64:
			v, err := strconv.ParseFloat(string(q.Value), 64)
			if err != nil {
				return nil, err
			}
			if err := row.Modify(r, e.tb.GetColumnIndex(q.Column), v); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("unknown column type")
		}
		if err := e.tb.SetRow(q.Key, r); err != nil {
			return nil, err
		}
		return okResponse, nil
	case query.GetQuery:
		r, err := e.tb.GetRow(q.Key)
		if err != nil {
			return nil, err
		}
		resp := []byte(nil)
		switch e.tb.GetColumnType(q.Column) {
		case column.Bool:
			v, err := row.Get[bool](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatBool(v))
		case column.Bytes:
			v, err := row.Get[[]byte](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = v
		case column.String:
			v, err := row.Get[string](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(v)
		case column.Int8:
			v, err := row.Get[int8](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatInt(int64(v), 10))
		case column.Int16:
			v, err := row.Get[int16](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatInt(int64(v), 10))
		case column.Int32:
			v, err := row.Get[int32](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatInt(int64(v), 10))
		case column.Int64:
			v, err := row.Get[int64](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatInt(v, 10))
		case column.Uint8:
			v, err := row.Get[uint8](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatUint(uint64(v), 10))
		case column.Uint16:
			v, err := row.Get[uint16](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatUint(uint64(v), 10))
		case column.Uint32:
			v, err := row.Get[uint32](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatUint(uint64(v), 10))
		case column.Uint64:
			v, err := row.Get[uint64](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatUint(v, 10))
		case column.Float32:
			v, err := row.Get[float32](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatFloat(float64(v), 'f', -1, 32))
		case column.Float64:
			v, err := row.Get[float64](r, e.tb.GetColumnIndex(q.Column))
			if err != nil {
				return nil, err
			}
			resp = []byte(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			return nil, errors.New("unknown column type")
		}
		return resp, nil
	case query.DelQuery:
		if err := e.tb.DeleteRow(q.Key); err != nil {
			return nil, err
		}
		return okResponse, nil
	case query.LockQuery:
		if ok := e.tb.Lock(q.Key); !ok {
			return failResponse, nil
		}
		return okResponse, nil
	case query.UnlockQuery:
		if ok := e.tb.Unlock(q.Key); !ok {
			return failResponse, nil
		}
		return okResponse, nil
	}
	return []byte{}, nil
}

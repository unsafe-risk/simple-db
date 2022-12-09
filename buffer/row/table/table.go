package table

import (
	"errors"
	"github.com/unsafe-risk/simple-db/buffer/row"

	"github.com/cornelk/hashmap"
	"github.com/dgraph-io/badger/v3"
)

type Table struct {
	conn   *badger.DB
	locker *hashmap.Map[string, struct{}]

	columns        []int
	columnIndexMap *hashmap.Map[string, int]
}

func New(path string) (*Table, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &Table{
		conn:   db,
		locker: hashmap.New[string, struct{}](),
	}, nil
}

func (t *Table) SetColumns(n []string, c []int) bool {
	if len(n) != len(c) {
		return false
	}
	t.columns = c
	t.columnIndexMap = hashmap.New[string, int]()
	for i := range n {
		t.columnIndexMap.Set(n[i], i)
	}
	return true
}

func (t *Table) GetRow(key string) (*row.Row, error) {
	r := row.New(t.columns...)
	if err := t.conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		if err := item.Value(func(val []byte) error {
			return r.SetBytes(val)
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return r, nil
}

func (t *Table) GetBlankRow() (*row.Row, error) {
	return row.NewBlank(t.columns...)
}

func (t *Table) SetRow(key string, r *row.Row) error {
	if !r.EqualColumnTypes(t.columns...) {
		return errors.New("column type mismatch")
	}
	if err := t.conn.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), r.GetBytes())
	}); err != nil {
		return err
	}
	return nil
}

func (t *Table) DeleteRow(key string) error {
	if err := t.conn.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	}); err != nil {
		return err
	}
	return nil
}

func (t *Table) Lock(key string) bool {
	if _, exist := t.locker.GetOrInsert(key, struct{}{}); exist {
		return false
	}
	return true
}

func (t *Table) Unlock(key string) bool {
	return t.locker.Del(key)
}

func (t *Table) GetColumnIndex(col string) int {
	if t.columnIndexMap == nil {
		return -1
	}
	if index, exist := t.columnIndexMap.Get(col); exist {
		return index
	}
	return -1
}

func (t *Table) Close() error {
	return t.conn.Close()
}

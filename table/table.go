package table

import (
	"errors"

	"github.com/cornelk/hashmap"
	"github.com/dgraph-io/badger/v3"
	"github.com/unsafe-risk/simple-db/row"
)

type Table struct {
	conn   *badger.DB
	locker *hashmap.Map[string, struct{}]

	columns []int
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

func (t *Table) SetColumns(c ...int) {
	t.columns = c
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

func (t *Table) Close() error {
	return t.conn.Close()
}

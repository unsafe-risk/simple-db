package query

import (
	"bytes"
	"errors"
)

var (
	separator = []byte(" ")

	setQuery    = []byte("set")
	getQuery    = []byte("get")
	delQuery    = []byte("del")
	lockQuery   = []byte("lock")
	unlockQuery = []byte("unlock")
)

func Parse(buf []byte) (Query, error) {
	bytes.Trim(buf, " \r\n\t")
	l := bytes.Split(buf, separator)

	if len(l) < 1 {
		return Query{}, errors.New("invalid query")
	}

	switch {
	case bytes.Equal(l[0], setQuery):
		if len(l) < 4 {
			return Query{}, errors.New("invalid set query")
		}
		return Query{
			Type:   SetQuery,
			Key:    string(l[1]),
			Column: string(l[2]),
			Value:  l[3],
		}, nil
	case bytes.Equal(l[0], getQuery):
		if len(l) < 3 {
			return Query{}, errors.New("invalid get query")
		}
		return Query{
			Type:   GetQuery,
			Key:    string(l[1]),
			Column: string(l[2]),
		}, nil
	case bytes.Equal(l[0], delQuery):
		if len(l) < 3 {
			return Query{}, errors.New("invalid del query")
		}
		return Query{
			Type:   DelQuery,
			Key:    string(l[1]),
			Column: string(l[2]),
		}, nil
	case bytes.Equal(l[0], lockQuery):
		if len(l) < 2 {
			return Query{}, errors.New("invalid lock query")
		}
		return Query{
			Type: LockQuery,
			Key:  string(l[1]),
		}, nil
	case bytes.Equal(l[0], unlockQuery):
		if len(l) < 2 {
			return Query{}, errors.New("invalid unlock query")
		}
		return Query{
			Type: UnlockQuery,
			Key:  string(l[1]),
		}, nil
	}
	return Query{}, errors.New("unknown query")
}

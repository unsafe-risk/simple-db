package query

const (
	SetQuery = iota
	GetQuery
	DelQuery
	LockQuery
	UnlockQuery
)

type Query struct {
	Type  int
	Key   string
	Value []byte
}

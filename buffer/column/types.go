package column

type Column interface {
	[]byte | string | float64 | float32 | int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | bool
}

const (
	Bytes = iota
	String
	Float64
	Float32
	Int64
	Int32
	Int16
	Int8
	Uint64
	Uint32
	Uint16
	Uint8
	Bool
)

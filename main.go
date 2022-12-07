package main

import (
	"fmt"

	"github.com/unsafe-risk/simple-db/buffer"
	"github.com/unsafe-risk/simple-db/row"
	"github.com/unsafe-risk/simple-db/row/column"
)

func main() {
	wb := buffer.NewWriteBuffer()
	wb.Write([]byte{})
	wb.WriteString("")
	wb.WriteFloat64(3.141592)
	wb.WriteBool(true)
	wb.WriteInt64(1234567890)

	r := row.NewRow([]int{column.Bytes, column.String, column.Float64, column.Bool, column.Int64})
	r.SetBytes(wb.Bytes())
	row.Change[int64](r, 4, 9876543210)
	fmt.Println(row.Get[int64](r, 4))
}

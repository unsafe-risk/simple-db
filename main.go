package main

import (
	"fmt"

	"github.com/unsafe-risk/simple-db/buffer"
	"github.com/unsafe-risk/simple-db/row"
	"github.com/unsafe-risk/simple-db/row/column"
)

func main() {
	wb := buffer.NewWriteBuffer()
	wb.WriteString("John")
	wb.WriteInt8(20)
	wb.WriteInt8(120)

	r := row.NewRow(column.String, column.Int8, column.Int8)
	r.SetBytes(wb.Bytes())
	bs, err := row.Modify(r, 0, "merak")
	if err != nil {
		panic(err)
	}
	r.SetBytes(bs)
	fmt.Println(row.Get[string](r, 0))
	fmt.Println(row.Get[int8](r, 1))
	fmt.Println(row.Get[int8](r, 2))
}

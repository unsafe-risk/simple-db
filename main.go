package main

import (
	"fmt"

	"github.com/unsafe-risk/simple-db/row"
	"github.com/unsafe-risk/simple-db/row/column"
	"github.com/unsafe-risk/simple-db/table"
)

func main() {
	persons, err := table.New("datas/persons")
	if err != nil {
		panic(err)
	}

	persons.SetColumns(column.String, column.Int8)

	r, err := persons.GetBlankRow()
	if err != nil {
		panic(err)
	}

	row.Modify(r, 0, "John Doe")
	row.Modify(r, 1, int8(42))

	if err := persons.SetRow([]byte("1"), r); err != nil {
		panic(err)
	}

	r, err = persons.GetRow([]byte("1"))
	if err != nil {
		panic(err)
	}

	name, err := row.Get[string](r, 0)
	if err != nil {
		panic(err)
	}

	age, err := row.Get[int8](r, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(name, age)

	if err := persons.DeleteRow([]byte("1")); err != nil {
		panic(err)
	}

	if err := persons.Close(); err != nil {
		panic(err)
	}
}

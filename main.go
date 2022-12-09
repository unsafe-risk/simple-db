package main

import (
	"fmt"
	"os"
	"os/signal"

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

	if ok := persons.Lock("1"); !ok {
		panic("lock failed")
	}

	r, err := persons.GetBlankRow()
	if err != nil {
		panic(err)
	}

	row.Modify(r, 0, "John Doe")
	row.Modify(r, 1, int8(42))

	if err := persons.SetRow("1", r); err != nil {
		panic(err)
	}

	r, err = persons.GetRow("1")
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

	if err := persons.DeleteRow("1"); err != nil {
		panic(err)
	}

	if ok := persons.Unlock("1"); !ok {
		panic("unlock failed")
	}

	if err := persons.Close(); err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}

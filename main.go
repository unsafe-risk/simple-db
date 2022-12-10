package main

import (
	"fmt"
	"github.com/unsafe-risk/simple-db/buffer/column"
	"github.com/unsafe-risk/simple-db/query/executor"
)

func main() {
	persons, err := executor.New()
	if err != nil {
		panic(err)
	}
	if err := persons.SetTable("persons", []string{"name", "age"}, []int{column.String, column.Int8}); err != nil {
		panic(err)
	}

	fmt.Println(persons.Run([]byte("set 1 name snowmerak")))
	fmt.Println(persons.Run([]byte("set 1 age 21")))
	fmt.Println(persons.Run([]byte("get 1 name")))
	fmt.Println(persons.Run([]byte("get 1 age")))

	if err := persons.Close(); err != nil {
		panic(err)
	}
}

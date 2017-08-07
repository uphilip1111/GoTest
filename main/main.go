package main

import (
	"dbtest/dbfunc"
	"dbtest/dborm"
	"dbtest/hello"
	"fmt"
)

func world() string {
	return "hello world!"
}

func main() {
	h := hello.Hello()
	w := world()
	dbfunc.QueryFunc()
	dborm.QueryFunc()
	fmt.Println(h)
	fmt.Println(w)
}

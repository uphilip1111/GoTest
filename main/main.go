package main

import (
	"dbtest/dbfunc"
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
	fmt.Println(h)
	fmt.Println(w)
}

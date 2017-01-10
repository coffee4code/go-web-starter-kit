package main

import (
	"github.com/go-martini/martini"
	"fmt"
)

func main() {
	fmt.Print("hello world")

	m := martini.Classic()

	m.Get("/", hello)
	m.Run()
}

func hello() string {
	return "hello world"
}
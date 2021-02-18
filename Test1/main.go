package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (U *user) notify() {
	fmt.Println("My name is", U.name, "and my email is", U.email)
}

type notification interface {
	notify()
}

func main() {
	test := user{"tiago", "tiago@ifro.edu.br"}
	test.notify()
}

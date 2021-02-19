package main

import (
	"fmt"
	"io/ioutil"
)

type send struct {
	Name  string
	Email string
	Phone string
}

func main() {

	doc, err := ioutil.ReadFile("../test.json")
	if err != nil {
		fmt.Println("erro")
	}
	fmt.Printf(string(doc) + "\n")

}

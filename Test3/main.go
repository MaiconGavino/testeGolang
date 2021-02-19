package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Send struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {

	doc, err := ioutil.ReadFile("../test.json")
	if err != nil {
		fmt.Println("erro")
	}
	dec := json.NewDecoder(strings.NewReader(string(doc)))
	for {
		var up []Send
		if err := dec.Decode(&up); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v \n", up)
	}
	//fmt.Printf(string(doc) + "\n")

}

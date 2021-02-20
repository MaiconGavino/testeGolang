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
		for _, value := range up {
			fmt.Println(value)
			ioutil.WriteFile("../Test3/exemplo.txt", []byte(value.Name), 0644)
		}
		//ioutil.WriteFile("../Test3/saida.txt", up, 0644)
	}
	//fmt.Printf(string(doc) + "\n")

}

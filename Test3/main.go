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
		for i, value := range up {
			fmt.Println(i, value)
			saida := []byte("Name:" + value.Name + "; Email:" + value.Email + "; Phone:" + value.Phone + "\n")
			err := ioutil.WriteFile("../Test3/exemplo.txt", saida, 0644)
			if err != nil {
				fmt.Println(err)
			}
		}
		//ioutil.WriteFile("../Test3/saida.txt", up, 0644)
	}
	//fmt.Printf(string(doc) + "\n")

}

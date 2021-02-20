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
		aux := "NAME:\t\t\t\t\t EMAIL:\t\t\t\t\t\t PHONE:"
		//var aux string
		for _, value := range up {
			teste := value
			aux = aux + "\n" + string([]byte(teste.Name)) + "\t\t" + string([]byte(teste.Email)) + "\t\t\t\t" + string([]byte(teste.Phone))
		}
		err := ioutil.WriteFile("../Test3/exemplo.txt", []byte(aux), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}

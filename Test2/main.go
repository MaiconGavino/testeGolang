package main

import (
	"fmt"
)

type Intern struct {
	Name     string
	Age      int
	Email    string
	Function string
}

func (user *Intern) company() {
	fmt.Println("Name: \t\t", user.Name, "\nAge:\t\t", user.Age, "\nEmail:\t\t", user.Email, "\nfunction:\t", user.Function, "\n")
}

type Employee struct {
	Name     string
	Age      int
	Email    string
	Function string
}

func (user *Employee) company() {
	fmt.Println("Name: \t\t", user.Name, "\nAge:\t\t", user.Age, "\nEmail:\t\t", user.Email, "\nfunction:\t", user.Function, "\n")
}

type Outsourced struct {
	Name     string
	Age      int
	Email    string
	Function string
}

func (user *Outsourced) company() {
	fmt.Println("Name: \t\t", user.Name, "\nAge:\t\t", user.Age, "\nEmail:\t\t", user.Email, "\nfunction:\t", user.Function, "\n")
}

type presentation interface {
	company()
}

func main() {
	junior := Intern{
		Name:     "Thiago",
		Age:      24,
		Email:    "thiago.gotec@ifro.edu.br",
		Function: "Junior Develop",
	}
	senior := Outsourced{
		Name:     "Vinicio",
		Age:      29,
		Email:    "vinicio.gotec@ifro.edu.br",
		Function: "Senior Develop",
	}
	pleno := Employee{
		Name:     "Bianca",
		Age:      26,
		Email:    "bianca.gotec@ifro.edu.br",
		Function: "Pleno Develop",
	}
	send(pleno, junior, senior)
}
func send(userE Employee, userI Intern, userO Outsourced) {
	userE.company()
	userI.company()
	userO.company()
}

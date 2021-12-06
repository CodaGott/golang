package main

import "fmt"

func main() {
	var a account

	var acc account

	acc.deposit()

	a = savings{typeOfAccount: ""}
	a.deposit()
	a.withdraw()
}

type account interface {
	deposit()
	withdraw()
}

type savings struct {
	typeOfAccount string
}

func (s savings) deposit(){
	s.typeOfAccount = "Savings struct Deposit method"
	fmt.Println(s)
}

func (s savings) withdraw(){
	s.typeOfAccount = "Savings struct Withdraw method"
	fmt.Println(s)
}

type test struct{
	name string
	age int32
}

func(t test){
	
}
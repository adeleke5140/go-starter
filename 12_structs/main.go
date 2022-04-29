package main

import (
	"fmt"
	"strconv"
)
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//pointer receiver because we are modifying the value in the struct instance
func (p *Person) hasBirthday(){
	p.age++ 
}

func (p *Person) getMarried(spouseLastName string){
	if p.gender =="m"{
		return
	} else {
		p.lastName = spouseLastName
	}
}



func main(){
	person1 := Person{firstName: "Kehinde", lastName: "Adeleke", city:"Lagos"}
	person2 := Person{"Olajesutofunmi", "Akinyemi", "Ekiti", "f", 21}

	person1.age = 21
	person1.gender = "m"
	person1.hasBirthday()

	person2.getMarried(person1.lastName)
	
	fmt.Println(person2)
	
}
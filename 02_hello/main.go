package main

import "fmt"

func main(){
	var age int = 21
	name, email := "Kehinde", "kehinde@hey.com"

	firstName:= "Adeleke"

	fmt.Println(name, age, firstName, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

}
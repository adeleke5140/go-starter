package main

import "fmt"

//a pointer is the location of a variable in Memory

func main(){
	  a := 5;
		b := &a

		c := *b
		*b = 10
		fmt.Println(a, b, *b, c)
}
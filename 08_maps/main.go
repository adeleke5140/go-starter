package main

import "fmt"

func main() {
	emails:= make(map[string]string)

	emails["Bob"]= "Bob@gmail.com"
	emails["ken"]= "Ken@hey.com"


	delete(emails, "Bob")
	fmt.Println(emails)
}
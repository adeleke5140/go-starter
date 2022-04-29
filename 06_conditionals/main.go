package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red"{
		fmt.Printf("%s is the best color in the world", color)
	}else {
		fmt.Printf("%s is still a great color\n", color)
	}


	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or not")
	}
}
package main

import "fmt"

func main() {
	ids := []int{33,76,54,2,11,23}

	for i, id := range ids{
		fmt.Printf("%d - ID: %d\n", i , id)
	}

	sum:= 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
}
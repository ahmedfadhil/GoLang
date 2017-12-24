package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	x := 5
	for {
		fmt.Println("Do stuff", x)
		x += 4
		if x > 25 {
			break
		}

	}
	a := 2
	for x := 4; a < 24; x += 3 {
		fmt.Println("This is a value:", x)
		a += 4
	}
}

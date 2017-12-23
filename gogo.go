package main

import (
	"fmt"
	"math"
)

func foo() {
	fmt.Println("The square root of 4 is:", math.Sqrt(4))
}

func add(x, y float64) float64 {
	return x + y

}

func multiple(a, b string) (string, string) {
	return a, b

}
func main() {

	x := 15
	a := &x // Memory address
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5 // Read memory address
	fmt.Println(x)

	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
	//
	//foo()
	//fmt.Println("A number from 1-100: ", rand.Intn(100))
	//var num1, num2 = 5.0, 4.6
	////var num2 float64 = 12.55
	//fmt.Println("Adding:", add(num1, num2))
	//
	//var w1, w2 = "Hey there", "brown cow"
	//
	//fmt.Print("Multiplication: ")
	//fmt.Println(multiple(w1, w2))
	//
	//var a int = 43
	//var b float64 = float64(a)
	//
	//var x = a
	//y := b
	//
	//fmt.Println(x)
	//fmt.Println(" ")
	//fmt.Println(y)

}

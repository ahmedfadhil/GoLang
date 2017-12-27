package main

import "fmt"

func main() {
	grades := make(map[string]float32)
	grades["timmy"] = 34
	grades["jimy"] = 45
	grades["watson"] = 56
	fmt.Println(grades)
	TimesGrade := grades["timmy"]
	fmt.Println(TimesGrade)
	delete(grades, "timmy")
	fmt.Println(grades)
	for k, v := range grades {
		fmt.Println(k, ":", v)

	}
}

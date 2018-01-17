package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Timmy"] = 23
	grades["Jimmy"] = 24
	grades["wat"] = 45

	fmt.Println(grades)
	TimeGrade := grades["Timmy"]
	fmt.Println(TimeGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for key, value := range grades {
		fmt.Println(key, value)
	}
}

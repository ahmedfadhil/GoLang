package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

//func index_handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hey there, Go is neat")
//}
//
//func about_handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "About page context here...")
//}
//
//func main() {
//	http.HandleFunc("/", index_handler)
//	http.HandleFunc("/", about_handler)
//	http.ListenAndServe(":8000", nil)
//}

type car struct {
	gas_pedal      uint64 // min 0 max 65535
	brake_pedal    uint64
	steering_wheel uint64
	top_speed_kmh  float64
}

func (c car) kmh() float64  {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
	
}

func (c car) mph() float64  {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)

}

func main() {
	a_car := car{gas_pedal: 65000,
	brake_pedal: 0,
	steering_wheel: 1256,
	top_speed_kmh: 225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

}

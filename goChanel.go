package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup
func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5

}

func main() {
	fooVal := make(chan int,10)
	for i := 0; i < 10; i++ {
		go foo(fooVal,i)
	}
	wg.Wait()
	close(fooVal)

	for item :=range fooVal{
		fmt.Println(item)
	}
//time.Sleep(time.Second*2)

	//go foo(fooVal, 5)
	//go foo(fooVal, 6)
	//
	//v1 := <-fooVal
	//v2 := <-fooVal
	//
	//fmt.Println(v1, v2)
}

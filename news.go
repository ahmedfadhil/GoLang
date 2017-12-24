package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news/sports/wp/2017/12/23/pressures-off-joe-flacco-and-john-harbaugh-with-ravens-nearing-playoff-berth/?hpid=hp_no-name_hp-in-the-news%3Apage%2Fin-the-news&utm_term=.5ebb5cc8928d")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()

}

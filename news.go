package main

import (
	"encoding/xml"
	"fmt"
	"sync"
	"net/http"
)

var wg sync.WaitGroup
var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

type NewsMap struct {
	keyword  string
	Location string
}

type SitemapIndex struct {
	Locations [] string `xml:"sitemap>loc"`
}

type News struct {
	Titles   [] string `xml:"url>news>title"`
	Keywords [] string `xml:"url>news>keywords"`
	Location [] string `xml:"url>loc"`
}

//type Location struct {
//	Loc string `xml:"loc"`
//}
//
//func (l Location) String() string {
//	return fmt.Sprintf(l.Loc)
//
//}

func newsRoutine(c chan News, location string) {
	defer wg.Done()
	var n News
	fmt.Printf("\n%s", Location)
	bytes := washPostXML
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n

}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news/sports/wp/2017/12/23/pressures-off-joe-flacco-and-john-harbaugh-with-ravens-nearing-playoff-berth/?hpid=hp_no-name_hp-in-the-news%3Apage%2Fin-the-news&utm_term=.5ebb5cc8928d")
	var s SitemapIndex

	news_map := make(map[string]NewsMap)
	resp.Body.Close()

	bytes := washPostXML
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
	queue := make(chan News, 30)
	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
		//fmt.Println(s.Locations)

	}
	wg.Wait()
	close(queue)
	for elem := range queue {
		for idx, _ := range elem.Titles {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Location[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n", idx)
		fmt.Println("\n", data.keyword)
		fmt.Println("\n", data.Location)

	}
}

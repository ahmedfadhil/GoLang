package main

import (
	"encoding/xml"
	"fmt"
)

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

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/news/sports/wp/2017/12/23/pressures-off-joe-flacco-and-john-harbaugh-with-ravens-nearing-playoff-berth/?hpid=hp_no-name_hp-in-the-news%3Apage%2Fin-the-news&utm_term=.5ebb5cc8928d")
	var s SitemapIndex
	var n News

	bytes := washPostXML
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)

		bytes := washPostXML
		xml.Unmarshal(bytes, &s)
		//fmt.Println(s.Locations)
	}
}

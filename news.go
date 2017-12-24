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
	Locations [] Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)

}

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/news/sports/wp/2017/12/23/pressures-off-joe-flacco-and-john-harbaugh-with-ravens-nearing-playoff-berth/?hpid=hp_no-name_hp-in-the-news%3Apage%2Fin-the-news&utm_term=.5ebb5cc8928d")
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)

}

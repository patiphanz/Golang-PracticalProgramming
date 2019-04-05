package main

import (
	"encoding/xml"
	"fmt"
)

var washPostXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/politics.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/business.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/going-out-guide.xml</loc>
	</sitemap>
`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	//resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}

package main

import "encoding/xml"
import "io/ioutil"
import "strings"

// Sitemap sitemap xml对象
type Sitemap struct {
	URLSet xml.Name `xml:"urlset"`
	URLs   []*URL   `xml:"url"`
}

// URL sitemap中的url对象
type URL struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

// ParseSitemap 解析sitemap文件生成Sitemap对象
func ParseSitemap(filename string) (*Sitemap, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	sitemap := &Sitemap{}
	err = xml.Unmarshal(data, sitemap)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

// SitemapToURLLabelName 根据sitemap的内容，产生url的label
func SitemapToURLLabelName(sitemap *Sitemap, domain string) []string {
	names := make([]string, 0, 0)
	if sitemap == nil {
		return names
	}

	domainLen := len(domain)
	for _, url := range sitemap.URLs {
		if strings.HasPrefix(url.Loc, domain) && strings.HasSuffix(url.Loc, "/") {
			names = append(names, url.Loc[domainLen:])
		}
	}

	return names
}

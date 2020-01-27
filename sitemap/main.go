package main

import (
	"flag"
	"fmt"
	"gophercises/link/parser"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/demo/cyoa", "provide url of website to map")
	flag.Parse()

	pages := getWeb(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func getWeb(urlFlag string) []string {
	resp, err := http.Get(urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(href(resp.Body, base), withPrefix(base))
}

func href(r io.Reader, base string) []string {
	links, _ := parser.Parse(r)
	var ret []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			ret = append(ret, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			ret = append(ret, link.Href)
		}
	}
	return ret
}

func filter(links []string, keepFnx func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFnx(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	} 

}

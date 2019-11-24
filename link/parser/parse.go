package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link --
type Link struct {
	Href string
	Text string
}

//Links -
type Links []Link

// Parse ---
func Parse(r io.Reader) (Links, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return getLinks(node), nil
}

func getLinks(node *html.Node) Links {
	var links Links
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				text := getText(node)
				links = append(links, Link{a.Val, text})
			}
		}
	}
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		l := getLinks(i)
		links = append(links, l...)
	}
	return links
}
func getText(node *html.Node) string {
	var text string
	if node.Type == html.TextNode && node.Data != "a" && node.Type != html.CommentNode {
		text = node.Data
	}
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		text += getText(i)
	}
	return strings.Trim(text, "\n")
}
